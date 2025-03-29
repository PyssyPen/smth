package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"nocon/delete"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type User struct {
	id    int
	email string
	logs  []logItem
}

type logItem struct {
	action    string
	timestamp time.Time
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}

	return out

}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

func saveUserInfo(user User, wg *sync.WaitGroup) error {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("writting file for user id: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644) // почитать про permission 0644 (read/write)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}

	wg.Done()

	return nil
}

func generateUser(count int, user chan User) {

	for i := 0; i < count; i++ {
		user <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail", i+1),
			logs:  generateLogs(500 + rand.Intn(500)),
		}
		time.Sleep(time.Millisecond * 10) // тут происходит блокировка
	}
	close(user) // нужно обязательно закрывать каналы после использования!
}

func main() {
	// для того, чтобы записать логи, надо сделать папку /VSC/GO/learn.udemy/2/logs
	rand.Seed(time.Now().Unix())

	t := time.Now()

	wg := &sync.WaitGroup{} //это некий счетчик,
	// который при наличии задач != 0 и будет уменьшаться до 0,
	// а когда это произойдет, переключится на следующую горутину

	users := make(chan User, 1000)
	go generateUser(100, users)

	for uSEr := range users {
		wg.Add(1)
		go saveUserInfo(uSEr, wg)
	}

	wg.Wait() // не пройдет дальше, пока != 0

	fmt.Println("Time elapsed:", time.Since(t).String())

	delete.Delete()
}

/*

func main() {
	go func() {
		time.Sleep(time.Second) // блокирует текущую горутину
		go fmt.Println("Конкурентный код с задержкой")
	}()

	go fmt.Println("Конкурентный код 1")
	go fmt.Println("Конкурентный код 2")
	go fmt.Println("Конкурентный код 3")

	time.Sleep(time.Second) // блокирует текущую горутину (в данном случае main)

	fmt.Println("Не конкурентный код")
}

*/
