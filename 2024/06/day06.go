package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	StartMarker = '^'
	Obstruction = '#'
)

type Position struct {
	Row, Col int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var moves = map[Direction][2]int{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
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

func findStart(field []string) (startPosition Position) {
	numRows := len(field)
	numCols := len(field[0])

	found := false
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if field[row][col] != StartMarker {
				continue
			}
			if found {
				log.Fatal("Start position must be a single location.")
			}
			startPosition = Position{row, col}
			found = true
		}
	}
	return startPosition
}

func patrol(
	field []string,
	startPosition Position,
) (positionToDirections map[Position][]Direction, canExit bool) {
	numRows := len(field)
	numCols := len(field[0])

	isInside := func(p Position) bool {
		return (0 <= p.Row && p.Row < numRows) && (0 <= p.Col && p.Col < numCols)
	}

	directions := []Direction{Up, Right, Down, Left}
	directionIndex := 0

	position := startPosition
	positionToDirections = make(map[Position][]Direction)
	for {
		direction := directions[directionIndex]
		if isInfinitePatrol(position, direction, positionToDirections) {
			return nil, false
		}

		positionToDirections[position] = append(positionToDirections[position], direction)

		move := moves[direction]
		nextPosition := Position{position.Row + move[0], position.Col + move[1]}
		if !isInside(nextPosition) {
			break
		}
		if field[nextPosition.Row][nextPosition.Col] == Obstruction {
			directionIndex = (directionIndex + 1) % len(directions)
			continue
		}
		position = nextPosition
	}

	return positionToDirections, true
}

func isInfinitePatrol(
	currentPosition Position,
	currentDirection Direction,
	positionToDirections map[Position][]Direction,
) bool {
	if directions, ok := positionToDirections[currentPosition]; ok {
		for _, d := range directions {
			// same position and direction encountered again
			if d == currentDirection {
				return true
			}
		}
	}
	return false
}

func main() {
	// const filename = "day06.example"
	const filename = "day06.input"

	field := parseInputFile(filename)

	// Part 1
	startPosition := findStart(field)
	positionToDirections, canExit := patrol(field, startPosition)
	if !canExit {
		log.Fatal("Cannot exit this field.")
	}
	fmt.Printf("number of patrolled positions: %d\n", len(positionToDirections))

	// Part 2
}
