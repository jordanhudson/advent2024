package main

import (
	"errors"
	"fmt"
)

var OffGrid = errors.New("went out of grid")

type StringMatrix [][]string

func (matrix StringMatrix) valAt(p Point) (string, error) {
	if p.y < 0 || p.y > len(matrix)-1 || p.x < 0 || p.x > len(matrix[p.y])-1 {
		return "", OffGrid
	}
	return matrix[p.y][p.x], nil
}

func (matrix StringMatrix) findFirst(str string) (Point, error) {
	for y, line := range matrix {
		for x, val := range line {
			if val == str {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, fmt.Errorf("string '%s' not found", str)
}

func readFileAsCharMatrix(filePath string) StringMatrix {
	slice := readFileAsStringSlice(filePath)
	matrix := make([][]string, len(slice))
	for y, line := range slice {
		matrix[y] = make([]string, len(line))
		for x, val := range line {
			matrix[y][x] = string(val)
		}
	}
	return matrix
}
