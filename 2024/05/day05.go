package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(filename string) (rules [][2]int, orders [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(data.String(), "\n\n")
	if len(sections) != 2 {
		log.Fatal("Input data must be two sections.")
	}

	rules = parseRules(strings.Fields(sections[0]))
	orders = parseOrders(strings.Fields(sections[1]))

	return rules, orders
}

func parseRules(lines []string) (rules [][2]int) {
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatal("Input data must be two parts.")
		}
		before, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}
		after, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, [2]int{before, after})
	}
	return rules
}

func parseOrders(lines []string) (orders [][]int) {
	for _, line := range lines {
		var order []int
		for _, page := range strings.Split(line, ",") {
			intPage, err := strconv.Atoi(page)
			if err != nil {
				log.Fatal(err)
			}
			order = append(order, intPage)
		}
		orders = append(orders, order)
	}
	return orders
}

func isCorrectlyOrdered(rules [][2]int, order []int) bool {
	pageToIndex := make(map[int]int)
	for i, page := range order {
		pageToIndex[page] = i
	}

	for _, rule := range rules {
		before := rule[0]
		after := rule[1]
		beforeIndex, ok := pageToIndex[before]
		if !ok {
			continue
		}
		afterIndex, ok := pageToIndex[after]
		if !ok {
			continue
		}
		if beforeIndex >= afterIndex {
			return false
		}
	}
	return true
}

func main() {
	// const filename = "day05.example"
	const filename = "day05.input"

	rules, orders := parseInputFile(filename)

	// Part 1
	result := 0
	for _, order := range orders {
		if isCorrectlyOrdered(rules, order) {
			result += order[len(order)/2]
		}
	}
	fmt.Printf("sum of the middle page number from correctly-ordered updates: %d\n", result)

	// Part 2
}
