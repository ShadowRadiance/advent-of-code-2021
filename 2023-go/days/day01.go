package days

import (
	"errors"
	"strconv"
	"strings"
)

type Day01 struct{}

const numberChars = "1234567890"

func extractNumericLiteral(line string, method func(string, string) int) (int, error) {
	idx := method(line, numberChars)
	if idx == -1 {
		return 0, errors.New("line does not contain any numeric literals")
	}
	number, _ := strconv.Atoi(string(line[idx]))
	return number, nil
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func (Day01) Part01(input string) string {
	lines := strings.Split(input, "\n")
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		first, err := extractNumericLiteral(line, strings.IndexAny)
		panicOnError(err)
		last, err := extractNumericLiteral(line, strings.LastIndexAny)
		panicOnError(err)
		numbers = append(numbers, first*10+last)
	}

	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}

	return strconv.Itoa(sum)
}

func (Day01) Part02(_ string) string { return "PENDING" }
