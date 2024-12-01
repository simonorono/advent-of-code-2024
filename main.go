package main

import "fmt"

type Day interface {
	part1() interface{}
	part2() interface{}
}

func main() {
	days := []Day{
		Day1{},
	}

	for idx, day := range days {
		fmt.Println("Day", idx+1)
		fmt.Println("  Part 1:", day.part1())
		fmt.Println("  Part 2:", day.part2())
	}
}
