package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input/day2.txt
var Day2Input string

type Day2 struct {
	input [][]int
}

func NewDay2() Day2 {
	lines := strings.Split(strings.TrimSpace(Day2Input), "\n")

	var input [][]int

	for _, line := range lines {
		split := strings.Split(line, " ")
		arr := make([]int, len(split))

		for i := 0; i < len(arr); i++ {
			arr[i], _ = strconv.Atoi(split[i])
		}

		input = append(input, arr)
	}

	return Day2{input}
}

func isDecreasing(line []int) bool {
	return line[0] > line[1]
}

func validateLine(line []int) bool {
	// will try to match the order of the first two elements for the rest of
	// the list
	decreasing := isDecreasing(line)

	for i := 0; i < len(line)-1; i++ {
		if decreasing && line[i] < line[i+1] {
			return false
		}

		if !decreasing && line[i+1] < line[i] {
			return false
		}

		diff := line[i] - line[i+1]

		if diff < 0 {
			diff *= -1
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	// did not find any errors
	return true
}

func (day Day2) part1() interface{} {
	result := 0

	for _, line := range day.input {
		if validateLine(line) {
			result++
		}
	}

	return result
}

func (day Day2) part2() interface{} {
	result := 0

outer:
	for _, line := range day.input {
		// original line valid? no need to check anything
		if validateLine(line) {
			result++
			continue
		}

		for i := range line {
			// create copy of line as to not mess with the original due to
			// golang's append semantics
			lineCopy := make([]int, len(line))
			copy(lineCopy, line)

			// create new slice sans the element at i
			minusElement := append(lineCopy[:i], lineCopy[i+1:]...)

			// any version of the line sans one element is valid? all line is
			// valid
			if validateLine(minusElement) {
				result++
				continue outer
			}
		}
	}

	return result
}
