package main

import (
	"flag"
	"fmt"
	"os"
)

type Day interface {
	part1() interface{}
	part2() interface{}
}

func executeDay(idx uint, day Day) {
	fmt.Println("Day", idx)
	fmt.Println("  Part 1:", day.part1())
	fmt.Println("  Part 2:", day.part2())
}

func main() {
	specificDay := flag.Uint("day", 0, "execute specific day")
	flag.Parse()

	days := []Day{
		NewDay1(),
		NewDay2(),
		Day3{},
	}

	if *specificDay > 0 && *specificDay > uint(len(days)) {
		fmt.Println("day doesn't exist")
		os.Exit(1)
	}

	if *specificDay > 0 && *specificDay <= uint(len(days)) {
		executeDay(*specificDay, days[*specificDay-1])
		return
	}

	for idx, day := range days {
		executeDay(uint(idx+1), day)
	}
}
