package main

import (
	"errors"
	"fmt"
)

func day6() {
	matrix := readFileAsCharMatrix("inputs/day6.txt")
	positions := make(map[Point]struct{})
	guardPos, _ := matrix.findFirst("^")
	direction := North
	for {
		positions[guardPos] = struct{}{}
		// look what's in front
		inFront := guardPos.getAdjacent(direction)
		valInFront, err := matrix.valAt(inFront)
		if errors.Is(err, OffGrid) {
			break // finished
		}
		// maybe turn
		if valInFront == "#" {
			direction = direction.turnRight90()
			inFront = guardPos.getAdjacent(direction)
		}
		// walk
		guardPos = inFront
	}
	fmt.Println(len(positions))
}
