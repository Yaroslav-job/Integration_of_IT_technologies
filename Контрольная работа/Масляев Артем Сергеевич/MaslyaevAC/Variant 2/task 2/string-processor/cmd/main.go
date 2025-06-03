package main

import (
	"string-processor/internal/generator"
	"string-processor/internal/filter"
)

func main() {
	strChan := make(chan string, generator.NumStrings)

	go generator.GenerateStrings(strChan)
	filter.FilterStrings(strChan)
}
