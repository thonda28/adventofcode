package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFile(filename string) (reports [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var report []int
		for _, data := range line {
			num, err := strconv.Atoi(data)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return reports
}

func isSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}
	if report[0] == report[len(report)-1] {
		return false
	}

	order := 1
	if report[0] > report[len(report)-1] {
		order = -1
	}

	for i := 1; i < len(report); i++ {
		diff := (report[i] - report[i-1]) * order
		if !(1 <= diff && diff <= 3) {
			return false
		}
	}
	return true
}

func countSafeReports(reports [][]int) (safeReportCount int) {
	for _, report := range reports {
		if isSafe(report) {
			safeReportCount++
		}
	}
	return safeReportCount
}

func countDampenedSafeReports(reports [][]int) (dampenedSafeReportCount int) {
	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			subReport := make([]int, 0, len(report)-1)
			subReport = append(subReport, report[:i]...)
			subReport = append(subReport, report[i+1:]...)

			if isSafe(subReport) {
				dampenedSafeReportCount++
				break
			}
		}
	}
	return dampenedSafeReportCount
}

func main() {
	// const filename = "day02.example"
	const filename = "day02.input"

	reports := parseInputFile(filename)

	// Part 1
	safeReportCount := countSafeReports(reports)
	fmt.Printf("number of safe reports: %d\n", safeReportCount)

	// Part 2
	safeReportWithDampenerCount := countDampenedSafeReports(reports)
	fmt.Printf("number of safe reports using dampener: %d\n", safeReportWithDampenerCount)
}
