package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct {
	list1 []int
	list2 []int
}

//go:embed input/day1.txt
var Day1Input string

func NewDay1() Day1 {
	lines := strings.Split(strings.TrimSpace(Day1Input), "\n")
	count := len(lines)

	list1 := make([]int, count)
	list2 := make([]int, count)

	for idx, value := range lines {
		parts := strings.Split(value, "   ")
		list1[idx], _ = strconv.Atoi(parts[0])
		list2[idx], _ = strconv.Atoi(parts[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return Day1{
		list1,
		list2,
	}
}

func (day Day1) part1() interface{} {
	lines := strings.Split(strings.TrimSpace(Day1Input), "\n")
	count := len(lines)

	result := 0

	for i := 0; i < count; i++ {
		diff := day.list1[i] - day.list2[i]

		if diff < 0 {
			diff *= -1
		}

		result += diff
	}

	return result
}

func countOccurrences(value int, slice []int) int {
	count := 0

	for _, v := range slice {
		if value == v {
			count++
		}
	}

	return count
}

func (day Day1) part2() interface{} {
	result := 0

	for _, v := range day.list1 {
		result += v * countOccurrences(v, day.list2)
	}

	return result
}
