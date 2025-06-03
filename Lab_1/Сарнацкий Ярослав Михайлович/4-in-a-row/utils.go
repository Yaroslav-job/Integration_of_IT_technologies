package main

import (
	"math/rand"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func random() int {
	return rand.Intn(7) + 1
}
