package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shadowradiance/advent-of-code/2023-go/aoc"
)

var scanner = bufio.NewScanner(os.Stdin)

func userInput(output string) string {
	fmt.Print(output)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var err error

	for _, solution := range aoc.Available() {
		aoc.RegisterDay(solution)
	}

	dayStr := userInput("Day [Enter for all]: ")

	if dayStr == "" {
		aoc.RunAllParts()
	} else {
		if day, err := strconv.Atoi(dayStr); err == nil {
			partStr := userInput("Part [Enter for all]: ")
			if partStr == "" {
				aoc.RunOnePart(day, 1)
				aoc.RunOnePart(day, 2)
			} else {
				if part, err := strconv.Atoi(partStr); err == nil {
					aoc.RunOnePart(day, part)
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done!")
}
