package main

import (
	"bufio"
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

func canSolve(answer int, terms []int) bool {
	operators := []string{"+", "*"}

	var canSolveHelper func(currentValue, answer int, terms []int, operator string) bool
	canSolveHelper = func(currentValue, answer int, terms []int, operator string) bool {
		if len(terms) == 0 {
			return currentValue == answer
		}
		switch operator {
		case "+":
			currentValue += terms[0]
		case "*":
			currentValue *= terms[0]
		default:
			log.Fatalf("Invalid operator %s", operator)
		}

		canSolve := false
		for _, op := range operators {
			canSolve = canSolve || canSolveHelper(currentValue, answer, terms[1:], op)
		}

		return canSolve
	}

	return canSolveHelper(terms[0], answer, terms[1:], "+") || canSolveHelper(terms[0], answer, terms[1:], "*")
}

func main() {
	// const filename = "day07.example"
	const filename = "day07.input"

	candidates := parseInputFile(filename)

	// Part 1
	totalCalibrationResult := 0
	for _, candidate := range candidates {
		if canSolve(candidate.Answer, candidate.Terms) {
			totalCalibrationResult += candidate.Answer
		}
	}
	fmt.Printf("totalCalibrationResult: %d\n", totalCalibrationResult)

	// Part 2
}
