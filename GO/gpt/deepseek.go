package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn *websocket.Conn
}

var clients = make(map[*Client]bool)
var mutex = &sync.Mutex{}

func main() {
	// Запускаем Python скрипт
	cmd := exec.Command("python3", "-u", "/home/pyssy/VSC/PY/deepseek/deepseek.py")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Kill()

	// Чтение вывода из Python
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			msg := scanner.Text()
			broadcast(msg)
		}
		if err := scanner.Err(); err != nil {
			log.Println("Python output error:", err)
		}
	}()

	// HTTP обработчики
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(w, r, stdin)
	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Chat</title>
    </head>
    <body>
        <div id="chat" style="height: 400px; overflow-y: scroll; border: 1px solid #ccc; padding: 10px; margin-bottom: 10px;"></div>
        <input type="text" id="input" placeholder="Type message..." style="width: 300px; padding: 5px;">
        <button onclick="sendMessage()">Send</button>

        <script>
            const ws = new WebSocket('ws://' + window.location.host + '/ws');
            const chat = document.getElementById('chat');

            ws.onmessage = function(event) {
                chat.innerHTML += '<div>' + event.data + '</div>';
                chat.scrollTop = chat.scrollHeight;
            };

            function sendMessage() {
                const input = document.getElementById('input');
                ws.send(input.value);
                input.value = '';
            }

            document.getElementById('input').addEventListener('keypress', function(e) {
                if (e.key === 'Enter') sendMessage();
            });
        </script>
    </body>
    </html>
    `)
}

func wsHandler(w http.ResponseWriter, r *http.Request, stdin io.WriteCloser) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade error:", err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn}
	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		stdin.Write([]byte(string(msg) + "\n"))
	}

	mutex.Lock()
	delete(clients, client)
	mutex.Unlock()
}

func broadcast(msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		err := client.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			client.conn.Close()
			delete(clients, client)
		}
	}
}
