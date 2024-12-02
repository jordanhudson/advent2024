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
	for _, line := range lines {
		re := regexp.MustCompile(`\d+`)
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
	for _, v := range listB {
		histogramOfB[v]++
	}
	totalSimilarityScore := 0
	for _, v := range listA {
		similarity := v * histogramOfB[v]
		totalSimilarityScore += similarity
	}
	fmt.Printf("total similarity: %d\n", totalSimilarityScore)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
