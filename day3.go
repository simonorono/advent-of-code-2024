package main

import (
	_ "embed"
	"regexp"
	"strconv"
)

const MulRegex = `mul\((\d{1,3}),(\d{1,3})\)`
const DoRegex = `do\(\)`
const DontRegex = `don't\(\)`

type Day3 struct{}

//go:embed input/day3.txt
var Day3Input string

func (day Day3) part1() interface{} {
	result := 0

	regex := regexp.MustCompile(MulRegex)

	matches := regex.FindAllStringSubmatch(Day3Input, -1)

	for _, v := range matches {
		left, _ := strconv.Atoi(v[1])
		right, _ := strconv.Atoi(v[2])
		result += left * right
	}

	return result
}

type Day3Range struct {
	from int
	to   int
}

func (day Day3) part2() interface{} {
	doRegex := regexp.MustCompile(DoRegex)
	dontRegex := regexp.MustCompile(DontRegex)

	// will be used to map the match results only at the beginning, don't care
	// about where they end
	onlyStart := func(array [][]int) []int {
		result := make([]int, len(array))

		for i, v := range array {
			result[i] = v[0]
		}

		return result
	}

	// checks if an int is inside an array of ints
	inArray := func(arr []int, el int) bool {
		for _, v := range arr {
			if el == v {
				return true
			}
		}

		return false
	}

	doIndices := onlyStart(doRegex.FindAllStringIndex(Day3Input, -1))
	dontIndices := onlyStart(dontRegex.FindAllStringIndex(Day3Input, -1))

	var ranges []Day3Range

	doing := true
	doingFrom := 0

	for i := 0; i < len(Day3Input); i++ {
		if doing && inArray(dontIndices, i) {
			doing = false
			ranges = append(ranges, Day3Range{doingFrom, i})
		}

		if !doing && inArray(doIndices, i) {
			doing = true
			doingFrom = i
		}
	}

	if doing {
		ranges = append(ranges, Day3Range{doingFrom, len(Day3Input) - 1})
	}

	regex := regexp.MustCompile(MulRegex)
	result := 0

	for _, r := range ranges {
		str := Day3Input[r.from : r.to+1]

		matches := regex.FindAllStringSubmatch(str, -1)

		for _, v := range matches {
			left, _ := strconv.Atoi(v[1])
			right, _ := strconv.Atoi(v[2])
			result += left * right
		}
	}

	return result
}
