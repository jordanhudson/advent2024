package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func day1() {
	lines := readFileAsStringSlice("inputs/day1.txt")
	var listA []int
	var listB []int
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		matches := re.FindAllString(line, 2)
		locationA, _ := strconv.Atoi(matches[0])
		locationB, _ := strconv.Atoi(matches[1])
		listA = append(listA, locationA)
		listB = append(listB, locationB)
	}
	sort.Ints(listA)
	sort.Ints(listB)

	// part 1
	totalDistance := 0
	for i := 0; i < len(listA); i++ {
		totalDistance += abs(listA[i] - listB[i])
	}
	fmt.Printf("total distance: %d\n", totalDistance)

	// part 2
	histogramOfB := make(map[int]int)
	for _, x := range listB {
		histogramOfB[x]++
	}
	totalSimilarityScore := 0
	for _, x := range listA {
		totalSimilarityScore += x * histogramOfB[x]
	}
	fmt.Printf("total similarity: %d\n", totalSimilarityScore)
}
