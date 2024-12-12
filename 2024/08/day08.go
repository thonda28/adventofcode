package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Position struct {
	Row, Col int
}

func parseInputFile(filename string) (field []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return field
}

func getAntennaPositions(field []string) (antennaPositions map[byte][]Position) {
	numRows := len(field)
	numCols := len(field[0])

	var isAntenna = func(b byte) bool {
		r := rune(b)
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}

	antennaPositions = make(map[byte][]Position)
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			antenna := field[row][col]
			if !isAntenna(antenna) {
				continue
			}

			position := Position{row, col}
			antennaPositions[antenna] = append(antennaPositions[antenna], position)
		}
	}

	return antennaPositions
}

func getAntinodes(
	field []string,
	positions []Position,
	isExtendedMode bool,
) (antinodePositions map[Position]struct{}) {
	numRows := len(field)
	numCols := len(field[0])

	isInside := func(row, col int) bool {
		return 0 <= row && row < numRows && 0 <= col && col < numCols
	}

	antinodePositions = make(map[Position]struct{})
	for i := 0; i < len(positions); i++ {
		posA := positions[i]
		for j := i + 1; j < len(positions); j++ {
			posB := positions[j]

			deltaRow := posA.Row - posB.Row
			deltaCol := posA.Col - posB.Col

			if isExtendedMode {
				// Part 2

				// antinode candidate on A’s side
				candidateA := posA
				for isInside(candidateA.Row, candidateA.Col) {
					antinodePositions[candidateA] = struct{}{}
					candidateA.Row += deltaRow
					candidateA.Col += deltaCol
				}

				// antinode candidate on B’s side
				candidateB := posB
				for isInside(candidateB.Row, candidateB.Col) {
					antinodePositions[candidateB] = struct{}{}
					candidateB.Row -= deltaRow
					candidateB.Col -= deltaCol
				}
			} else {
				// Part 1

				// antinode candidate on A’s side
				candidateA := Position{posA.Row + deltaRow, posA.Col + deltaCol}
				if isInside(candidateA.Row, candidateA.Col) {
					antinodePositions[candidateA] = struct{}{}
				}

				// antinode candidate on B’s side
				candidateB := Position{posB.Row - deltaRow, posB.Col - deltaCol}
				if isInside(candidateB.Row, candidateB.Col) {
					antinodePositions[candidateB] = struct{}{}
				}
			}
		}
	}
	return antinodePositions
}

func getAllAntinodes(
	field []string,
	antennaPositions map[byte][]Position,
	isExtendedMode bool,
) (allAntinodePositions map[Position]struct{}) {
	allAntinodePositions = make(map[Position]struct{})
	for _, positions := range antennaPositions {
		allAntinodePositions = union(allAntinodePositions, getAntinodes(field, positions, isExtendedMode))
	}
	return allAntinodePositions
}

func union(set1, set2 map[Position]struct{}) map[Position]struct{} {
	result := make(map[Position]struct{})

	for key := range set1 {
		result[key] = struct{}{}
	}

	for key := range set2 {
		result[key] = struct{}{}
	}

	return result
}

func main() {
	// const filename = "day08.example"
	const filename = "day08.input"

	field := parseInputFile(filename)

	// Part 1
	antennaPositions := getAntennaPositions(field)
	allAntinodes := getAllAntinodes(field, antennaPositions, false)
	fmt.Printf("the number of antinodes: %d\n", len(allAntinodes))

	// Part 2
	allExtendedAntinodes := getAllAntinodes(field, antennaPositions, true)
	fmt.Printf("the number of extended antinodes: %d\n", len(allExtendedAntinodes))
}
