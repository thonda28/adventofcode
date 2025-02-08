package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	StartMarker   = "^"
	Obstruction   = "#"
	NoObstruction = "."
)

type Position struct {
	Row, Col int
}

type Direction int

type PositionDirection struct {
	pos Position
	dir Direction
}

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
			if string(field[row][col]) != StartMarker {
				continue
			}
			if found {
				log.Fatal("Start position must be a single location.")
			}
			startPosition = Position{row, col}
			found = true
		}
	}

	if !found {
		log.Fatal("Start position does not exist.")
	}
	return startPosition
}

func patrol(
	field []string,
	startPosition Position,
) (visited map[Position]struct{}, canExit bool) {
	numRows := len(field)
	numCols := len(field[0])

	isInside := func(p Position) bool {
		return (0 <= p.Row && p.Row < numRows) && (0 <= p.Col && p.Col < numCols)
	}

	directions := []Direction{Up, Right, Down, Left}
	directionIndex := 0

	position := startPosition
	visited = make(map[Position]struct{})
	positionDirectionHistory := make(map[PositionDirection]struct{})
	for {
		direction := directions[directionIndex]
		if isInfinitePatrol(position, direction, positionDirectionHistory) {
			return nil, false
		}
		visited[position] = struct{}{}
		positionDirectionHistory[PositionDirection{position, direction}] = struct{}{}

		move := moves[direction]
		nextPosition := Position{position.Row + move[0], position.Col + move[1]}
		if !isInside(nextPosition) {
			break
		}
		if string(field[nextPosition.Row][nextPosition.Col]) == Obstruction {
			directionIndex = (directionIndex + 1) % len(directions)
			continue
		}
		position = nextPosition
	}

	return visited, true
}

func isInfinitePatrol(
	currentPosition Position,
	currentDirection Direction,
	positionDirectionHistory map[PositionDirection]struct{},
) bool {
	posDir := PositionDirection{currentPosition, currentDirection}
	// same position and direction encountered again
	_, ok := positionDirectionHistory[posDir]
	return ok
}

func countStuckableObstruction(field []string, startPosition Position) (stuckableObstructionCount int) {
	numRows := len(field)
	numCols := len(field[0])

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if row == startPosition.Row && col == startPosition.Col {
				continue
			}
			if string(field[row][col]) == Obstruction {
				continue
			}

			// place an obstruction temporarily
			field[row] = field[row][:col] + Obstruction + field[row][col+1:]

			if _, canExit := patrol(field, startPosition); !canExit {
				stuckableObstructionCount++
			}

			// remove the temporarily placed obstruction
			field[row] = field[row][:col] + NoObstruction + field[row][col+1:]
		}
	}

	return stuckableObstructionCount
}

func main() {
	// const filename = "day06.example"
	const filename = "day06.input"

	field := parseInputFile(filename)

	// Part 1
	startPosition := findStart(field)
	visitedPositions, canExit := patrol(field, startPosition)
	if !canExit {
		log.Fatal("Cannot exit this field.")
	}
	fmt.Printf("number of patrolled positions: %d\n", len(visitedPositions))

	// Part 2
	stuckableObstructionCount := countStuckableObstruction(field, startPosition)
	fmt.Printf("number of positions where placing an obstruction can make the guard stuck: %d\n", stuckableObstructionCount)
}
