package server

import (
	"net/http"

	"test-web/handlers"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ChatHandler)
	r.HandleFunc("/chat/{id}", handlers.ChatMessagesHandler)

	// Добавляем обработку статики
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
