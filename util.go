package main

import (
	"os"
	"strconv"
	"strings"
)

func readFileAsString(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func readFileAsStringSlice(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func stringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func removeElementAt(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}
	copiedSlice := append([]int(nil), slice...)
	return append(copiedSlice[:i], copiedSlice[i+1:]...)
}

func countOccurrences(slice []string, target string) int {
	count := 0
	for _, str := range slice {
		if str == target {
			count++
		}
	}
	return count
}

func indexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
