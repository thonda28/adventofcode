package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type MulOperation struct {
	operand1 int
	operand2 int
}

func (mulop MulOperation) multiply() int {
	return mulop.operand1 * mulop.operand2
}

func parseInputFile(filename string) (instruction string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return instruction
}

func extractMulOperations(instruction string) (mulOperations []MulOperation) {
	mul_pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re, err := regexp.Compile(mul_pattern)
	if err != nil {
		log.Fatal(err)
	}

	// https://pkg.go.dev/github.com/shogo82148/std/regexp#Regexp.FindAllStringSubmatch
	matches := re.FindAllStringSubmatch(instruction, -1)

	// exmaple of matches:
	// matches = [][]string{
	// 	{ "mul(2,4)",  "2", "4"},
	// 	{ "mul(5,5)",  "5", "5"},
	// 	{"mul(11,8)", "11", "8"},
	// 	{ "mul(8,5)",  "8", "5"},
	// }

	for _, match := range matches {
		mulOperand1, _ := strconv.Atoi(match[1])
		mulOperand2, _ := strconv.Atoi(match[2])
		mulOperation := MulOperation{mulOperand1, mulOperand2}
		mulOperations = append(mulOperations, mulOperation)
	}
	return mulOperations
}

func main() {
	// const filename = "day03.example"
	const filename = "day03.input"

	instruction := parseInputFile(filename)

	// Part 1
	mulOperations := extractMulOperations(instruction)
	result := 0
	for _, mulOperation := range mulOperations {
		result += mulOperation.multiply()
	}
	fmt.Printf("sum of the results of all multiplications: %d\n", result)

	// Part 2
}
