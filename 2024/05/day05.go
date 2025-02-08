package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}

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

func sortByRules(rules [][2]int, order []int) (sortedOrder []int) {
	pagesInOrder := make(map[int]struct{})
	for _, page := range order {
		pagesInOrder[page] = struct{}{}
	}

	beforeToAfters := make(map[int][]int)
	beforePageCount := make(map[int]int)
	for _, rule := range rules {
		before := rule[0]
		if _, ok := pagesInOrder[before]; !ok {
			continue
		}
		after := rule[1]
		if _, ok := pagesInOrder[after]; !ok {
			continue
		}
		beforeToAfters[before] = append(beforeToAfters[before], after)
		beforePageCount[after]++
	}

	var independentPages Stack[int]
	for _, page := range order {
		if beforePageCount[page] == 0 {
			independentPages.Push(page)
		}
	}

	for len(independentPages.data) > 0 {
		beforePage, _ := independentPages.Pop()
		for _, afterPage := range beforeToAfters[beforePage] {
			beforePageCount[afterPage]--
			if beforePageCount[afterPage] == 0 {
				independentPages.Push(afterPage)
			}
		}
		sortedOrder = append(sortedOrder, beforePage)
	}
	return sortedOrder
}

// Constraints:
// 1. The result of the Topological Sort is uniquely determined
//    1-1. The graph is a Directed Acyclic Graph (DAG)
//    1-2. At each step, there is only one vertex with an in-degree of 0
// 2. The length of the order is always odd
// 3. All elements in the order are unique

func main() {
	// const filename = "day05.example"
	const filename = "day05.input"

	rules, orders := parseInputFile(filename)

	// Part 1
	sumOfCorrectOrders := 0
	for _, order := range orders {
		if isCorrectlyOrdered(rules, order) {
			sumOfCorrectOrders += order[len(order)/2]
		}
	}
	fmt.Printf("sum of the middle page number from correctly-ordered updates: %d\n", sumOfCorrectOrders)

	// Part 2
	sumOfIncorrectOrders := 0
	for _, order := range orders {
		sortedOrder := sortByRules(rules, order)
		if !reflect.DeepEqual(order, sortedOrder) {
			sumOfIncorrectOrders += sortedOrder[len(sortedOrder)/2]
		}
	}
	fmt.Printf("sum of the middle page number from incorrectly-ordered updates: %d\n", sumOfIncorrectOrders)
}
