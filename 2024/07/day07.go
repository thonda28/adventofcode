package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Candidate struct {
	Answer int
	Terms  []int
}

func parseInputFile(filename string) (candidates []Candidate) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ":")
		answer, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		terms := []int{}
		for _, term := range strings.Fields(strings.TrimSpace(fields[1])) {
			term, err := strconv.Atoi(term)
			if err != nil {
				log.Fatal(err)
			}
			terms = append(terms, term)
		}

		candidates = append(candidates, Candidate{answer, terms})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return candidates
}

func canSolve(answer int, terms []int, operators []string) bool {
	if len(terms) == 0 {
		return false
	}

	var canSolveHelper func(currentValue, answer int, terms []int) bool
	canSolveHelper = func(currentValue, answer int, terms []int) bool {
		if len(terms) == 0 {
			return currentValue == answer
		}

		for _, op := range operators {
			nextValue, err := calculate(currentValue, terms[0], op)
			if err != nil {
				log.Fatal(err)
			}

			if canSolveHelper(nextValue, answer, terms[1:]) {
				return true
			}
		}
		return false
	}

	return canSolveHelper(terms[0], answer, terms[1:])
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	case "||":
		concatnated := strconv.Itoa(a) + strconv.Itoa(b)
		concatnatedInt, err := strconv.Atoi(concatnated)
		if err != nil {
			return 0, err
		}
		return concatnatedInt, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	// const filename = "day07.example"
	const filename = "day07.input"

	candidates := parseInputFile(filename)

	// Part 1
	twoAvailableOperators := []string{"+", "*"}
	totalCalibrationResult := 0
	for _, candidate := range candidates {
		if canSolve(candidate.Answer, candidate.Terms, twoAvailableOperators) {
			totalCalibrationResult += candidate.Answer
		}
	}
	fmt.Printf("totalCalibrationResult: %d\n", totalCalibrationResult)

	// Part 2
	threeAvailableOperators := []string{"+", "*", "||"}
	newTotalCalibrationResult := 0
	for _, candidate := range candidates {
		if canSolve(candidate.Answer, candidate.Terms, threeAvailableOperators) {
			newTotalCalibrationResult += candidate.Answer
		}
	}
	fmt.Printf("newTotalCalibrationResult: %d\n", newTotalCalibrationResult)
}
