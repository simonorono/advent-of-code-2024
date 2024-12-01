package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

const YEAR = 2015

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func downloadDay(day int) []byte {
	client := http.Client{}
	cookie := fmt.Sprintf("session=%s", os.Args[1])
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", YEAR, day)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("COOKIE", cookie)

	response, err := client.Do(request)
	check(err)

	bytes, err := io.ReadAll(response.Body)
	check(err)

	return bytes
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Cookie argument is required")
		os.Exit(1)
	}

	today := time.Now()

	if today.Year() == YEAR && int(today.Month()) != 12 {
		fmt.Println("AOC has not started")
		os.Exit(1)
	}

	var maxDay = 25

	if today.Year() == YEAR {
		maxDay = int(math.Min(float64(today.Day()), float64(maxDay)))
	}

	_ = os.Mkdir("input", 0777)

	for day := 1; day <= maxDay; day += 1 {
		fmt.Println("Downloading day", day)

		file, _ := os.Create(fmt.Sprintf("input/day%d.txt", day))
		_, _ = file.Write(downloadDay(day))
	}
}
