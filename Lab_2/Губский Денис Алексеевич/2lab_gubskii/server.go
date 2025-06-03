package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type GameServer struct {
	Logic *GameLogic
}

func NewGameServer() *GameServer {
	return &GameServer{Logic: NewGameLogic()}
}

func (gs *GameServer) Serve() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(gs.Logic.Game)
		} else if r.Method == http.MethodPost {
			var req struct{ Direction string }
			json.NewDecoder(r.Body).Decode(&req)
			gs.Logic.Move(req.Direction)
			json.NewEncoder(w).Encode(gs.Logic.Game)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	server := NewGameServer()
	server.Serve()
}
