package main

import (
	"os"
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
