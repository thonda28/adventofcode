package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputFile(filename string) (something any) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		fmt.Println(line)

		// data processing
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return something
}

func main() {
	const filename = "dayxx.example"
	// const filename = "dayxx.input"

	something := parseInputFile(filename)
	fmt.Println(something)

	// Part 1

	// Part 2
}
