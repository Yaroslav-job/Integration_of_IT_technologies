package web

import (
        "log"
        "net/http"
        "strconv"
)

// StartWebServer starts the web server for the Sudoku game
func StartWebServer(port int) {
        // Set up static file server
        fs := http.FileServer(http.Dir("web/static"))
        http.Handle("/", fs)
        
        // API endpoints
        http.HandleFunc("/api/game/new", handleNewGame)
        http.HandleFunc("/api/game/move", handleMove)
        
        // Start server
        log.Printf("Starting Sudoku web server on port %d...\n", port)
        portStr := ":" + strconv.Itoa(port)
        log.Printf("Access the game at http://localhost%s\n", portStr)
        log.Fatal(http.ListenAndServe("0.0.0.0"+portStr, nil))
}