package main

import (
	"fmt"
	"regexp"
)

type Rule struct {
	left  int
	right int
}

type Update struct {
	vals []int
	str  string
}

func day5() {
	re := regexp.MustCompile(`\d+`)

	rules := make([]Rule, 0)
	for _, ruleStr := range readFileAsStringSlice("inputs/day5rules.txt") {
		ruleStrs := re.FindAllString(ruleStr, -1)
		ruleNums := stringsToInts(ruleStrs)
		rules = append(rules, Rule{ruleNums[0], ruleNums[1]})
	}

	updates := make([]Update, 0)
	for _, updateStr := range readFileAsStringSlice("inputs/day5updates.txt") {
		updateStrs := re.FindAllString(updateStr, -1)
		vals := stringsToInts(updateStrs)
		updates = append(updates, Update{vals, updateStr})
	}

	correctlyOrderedTotal := 0
	incorrectUpdates := make([]Update, 0)
	for _, update := range updates {
		pass := true
		for _, rule := range rules {
			if !isUpdatePassesRule(update, rule) {
				pass = false
				break
			}
		}
		if pass {
			middlePage := update.vals[len(update.vals)/2]
			correctlyOrderedTotal += middlePage
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	fmt.Println(correctlyOrderedTotal)

	incorrectlyOrderedTotal := 0
	for _, update := range incorrectUpdates {
		fix(update, rules)
		middlePage := update.vals[len(update.vals)/2]
		incorrectlyOrderedTotal += middlePage
	}
	fmt.Println(incorrectlyOrderedTotal)
}

func isUpdatePassesRule(update Update, rule Rule) bool {
	xIndex := indexOf(update.vals, rule.left)
	yIndex := indexOf(update.vals, rule.right)
	return xIndex == -1 || yIndex == -1 || xIndex < yIndex
}

func fix(update Update, rules []Rule) {
	vals := update.vals
	fixed := false
	for !fixed {
		fixed = true
		for _, rule := range rules {
			if !isUpdatePassesRule(update, rule) {
				i := indexOf(vals, rule.left)
				j := indexOf(vals, rule.right)
				vals[i], vals[j] = vals[j], vals[i]
				fixed = false
				break
			}
		}
	}
}
