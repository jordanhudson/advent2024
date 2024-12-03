package main

import (
	"fmt"
	"regexp"
)

func day2() {
	lines := readFileAsStringSlice("inputs/day2.txt")
	re := regexp.MustCompile(`\d+`)
	safeCount := 0
	for _, line := range lines {
		report := stringsToInts(re.FindAllString(line, -1))
		for i := -1; i < len(report); i++ {
			moddedReport := removeElementAt(report, i) // passing -1 doesn't remove anything
			if isReportSafe(moddedReport) {
				safeCount++
				break
			}
		}
	}
	fmt.Printf("Safe count: %d\n", safeCount)
}

func isReportSafe(report []int) bool {
	Inc := 1
	Dec := -1
	direction := 0
	for i := 1; i < len(report); i++ {
		levelA := report[i-1]
		levelB := report[i]
		if levelA == levelB {
			return false
		}
		if levelA < levelB {
			if direction == Dec {
				return false
			}
			if levelB-levelA > 3 {
				return false
			}
			direction = Inc
		} else {
			if direction == Inc {
				return false
			}
			if levelA-levelB > 3 {
				return false
			}
			direction = Dec
		}
	}
	return true
}
