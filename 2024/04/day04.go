package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInputFile(filename string) (texts []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		texts = append(texts, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return texts
}

func isFoundInDirection(
	texts []string,
	word string,
	startRow,
	startCol int,
	direction [2]int,
) (isFound bool) {
	numRows := len(texts)
	numCols := len(texts[0])

	row := startRow
	col := startCol
	wordIndex := 0
	for {
		isInside := 0 <= row && row < numRows && 0 <= col && col < numCols
		if !isInside {
			return false
		}
		if texts[row][col] != word[wordIndex] {
			return false
		}
		if wordIndex == len(word)-1 {
			return true
		}

		row += direction[0]
		col += direction[1]
		wordIndex += 1
	}
}

func countWord(texts []string, word string) (wordCount int) {
	numRows := len(texts)
	numCols := len(texts[0])

	directions := [][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			for _, direction := range directions {
				if isFoundInDirection(texts, word, row, col, direction) {
					wordCount++
				}
			}
		}
	}
	return wordCount
}

func main() {
	// const filename = "day04.example" // size: 10x10
	const filename = "day04.input" // size: 140x140

	texts := parseInputFile(filename)

	// Part 1
	word := "XMAS"
	wordCount := countWord(texts, word)
	fmt.Printf("number of \"%s\": %d\n", word, wordCount)

	// Part 2
}
