package main

import (
	"4-in-a-row/server" // Импортируем пакет server
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/start_game", server.StartGame)
	http.HandleFunc("/move", server.HandleMove)
	http.HandleFunc("/reset", server.HandleReset)
	http.HandleFunc("/start", server.HandleStart) // Обработчик для выбора типа игры
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.Execute(w, nil)
}
