package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInputFile(filename string) (leftList, rightList []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		l, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		leftList = append(leftList, l)

		r, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		rightList = append(rightList, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return leftList, rightList
}

func calcTotalDistance(leftList, rightList []int) (totalDistance int) {
	sort.Ints(leftList)
	sort.Ints(rightList)

	if len(leftList) != len(rightList) {
		log.Fatal("Length of left and right lists must be equal")
	}

	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}
	return totalDistance
}

func calcSimilarityScore(leftList, rightList []int) (similarityScore int) {
	var rightMap = make(map[int]int)
	for _, r := range rightList {
		rightMap[r]++
	}
	for _, l := range leftList {
		similarityScore += l * rightMap[l]
	}
	return similarityScore
}

func main() {
	// const filename = "day01.example"
	const filename = "day01.input"

	leftList, rightList := parseInputFile(filename)

	// Part 1
	totalDistance := calcTotalDistance(leftList, rightList)
	fmt.Printf("totalDistance: %d\n", totalDistance)

	// Part 2
	similarityScore := calcSimilarityScore(leftList, rightList)
	fmt.Printf("similarityScore: %d\n", similarityScore)
}
