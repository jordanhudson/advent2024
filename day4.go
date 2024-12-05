package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	Northeast
	East
	Southeast
	South
	Southwest
	West
	Northwest
)

func day4() {
	matrix := readFileAsCharMatrix("inputs/day4.txt")

	xmasCount1 := 0
	for y, line := range matrix {
		for x, val := range line {
			if val == "X" {
				views := []string{
					getView(matrix, x, y, North, 3),
					getView(matrix, x, y, Northeast, 3),
					getView(matrix, x, y, East, 3),
					getView(matrix, x, y, Southeast, 3),
					getView(matrix, x, y, South, 3),
					getView(matrix, x, y, Southwest, 3),
					getView(matrix, x, y, West, 3),
					getView(matrix, x, y, Northwest, 3),
				}
				xmasCount1 += countOccurrences(views, "MAS")
			}
		}
	}
	fmt.Println(xmasCount1)

	xmasCount2 := 0
	for y, line := range matrix {
		for x, val := range line {
			if val == "A" {
				corners :=
					getView(matrix, x, y, Northwest, 1) +
						getView(matrix, x, y, Northeast, 1) +
						getView(matrix, x, y, Southwest, 1) +
						getView(matrix, x, y, Southeast, 1)
				valid := []string{
					"MSMS", "MMSS", "SSMM", "SMSM",
				}
				xmasCount2 += countOccurrences(valid, corners)
			}
		}
	}
	fmt.Println(xmasCount2)

}

func getView(matrix [][]string, originX int, originY int, direction Direction, distance int) string {
	ret := ""
	for i := 1; i <= distance; i++ {
		y := originY
		x := originX
		switch direction {
		case North, Northeast, Northwest:
			y = originY - i
		case South, Southeast, Southwest:
			y = originY + i
		default:
		}
		switch direction {
		case East, Northeast, Southeast:
			x = originX + i
		case West, Northwest, Southwest:
			x = originX - i
		default:
		}
		if y < 0 || y > len(matrix)-1 {
			break
		}
		if x < 0 || x > len(matrix[y])-1 {
			break
		}
		ret += matrix[y][x]
	}
	return ret
}
