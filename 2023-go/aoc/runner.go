package aoc

import (
	"fmt"
	"github.com/shadowradiance/advent-of-code/2023-go/days"
	"os"
	"path"
)

var (
	registeredDays = make([]days.DayInterface, 0, 25)
	parts          = [2]func(days.DayInterface, string) string{days.DayInterface.Part01, days.DayInterface.Part02}
)

func RegisterDay(day days.DayInterface) {
	registeredDays = append(registeredDays, day)
}

func RunAllParts() {
	for idx := range registeredDays {
		RunOnePart(idx+1, 1)
		RunOnePart(idx+1, 2)
	}
}

func RunOnePart(day int, part int) {
	fmt.Println("Running day", day, "part", part)

	operation := getPart(part)
	result := operation(getDay(day), data(day))
	fmt.Println(result)
}

func getPart(part int) func(days.DayInterface, string) string {
	return parts[part-1]
}

func getDay(day int) days.DayInterface {
	return registeredDays[day-1]
}

func data(day int) string {
	// load input file based on day and part
	filename := fmt.Sprintf("day%02d.txt", day)
	if dat, err := os.ReadFile(path.Join("data", filename)); err != nil {
		panic(err)
	} else {
		return string(dat)
	}
}
