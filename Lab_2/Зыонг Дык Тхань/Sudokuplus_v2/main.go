package main

import (
	"flag"
	"fmt"
	"os"
	"sudokuplus/terminal"
	"sudokuplus/web"
)

func main() {
	// Command-line flags
	webMode := flag.Bool("web", false, "Start the game in web mode")
	port := flag.Int("port", 5000, "Port for web server")

	// Parse command-line arguments
	flag.Parse()

	if *webMode {
		// Start web version
		fmt.Println("Starting Sudoku Plus in web mode...")
		web.StartWebServer(*port)
	} else {
		// Start terminal version
		fmt.Println("Starting Sudoku Plus in terminal mode...")
		terminal.StartGame()
	}

	// Exit gracefully
	fmt.Println("Exiting Sudoku Plus. Goodbye!")
	os.Exit(0)
}
