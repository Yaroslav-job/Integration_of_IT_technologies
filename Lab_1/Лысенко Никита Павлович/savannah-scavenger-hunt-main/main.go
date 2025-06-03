package main

import (
	"math/rand"
	"time"

	"C/savannah-scavenger-hunt-main/internal"
)

func init() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func main() {
	internal.Execute()
}
