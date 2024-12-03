package main

import (
	"fmt"
	"regexp"
)

func day3() {
	reExpr := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	reMulOperands := regexp.MustCompile(`\d+`)
	input := readFileAsString("inputs/day3.txt")
	exprs := reExpr.FindAllString(input, -1)
	total := 0
	mulEnabled := true
	for _, expr := range exprs {
		cmd := expr[:3]
		if cmd == "mul" && mulEnabled {
			operands := reMulOperands.FindAllString(expr, -1)
			intOperands := stringsToInts(operands)
			total += intOperands[0] * intOperands[1]
		}
		if cmd == "don" {
			mulEnabled = false
		}
		if cmd == "do(" {
			mulEnabled = true
		}
	}
	fmt.Println(total)
}
