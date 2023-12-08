package days

import (
	"errors"
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"strconv"
	"strings"
)

type Day01 struct{}

var englishNumbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func englishNumber(input string) (int, bool) {
	number, exists := englishNumbers[input]
	return number, exists
}

func extractNumericLiteral(line string, method func(string, string) int) (int, error) {
	idx := method(line, util.NumberChars)
	if idx == -1 {
		return 0, errors.New("line does not contain any numeric literals")
	}
	number, _ := strconv.Atoi(string(line[idx]))
	return number, nil
}

func tryConvertNumber(s string) (int, error) {
	if len(s) == 1 {
		if number, err := strconv.Atoi(s); err == nil {
			return number, nil
		}
	} else {
		if number, ok := englishNumber(s); ok {
			return number, nil
		}
	}
	return 0, errors.New(s + " is not convertible to a number")
}

func extractFirstNumber(line string) (int, error) {
	for firstPosition := 0; firstPosition < len(line); firstPosition++ {
		for lastPosition := firstPosition + 1; lastPosition <= len(line); lastPosition++ {
			if number, err := tryConvertNumber(line[firstPosition:lastPosition]); err == nil {
				return number, nil
			}
		}
	}
	return 0, errors.New(line + " does not contain any numbers")
}

func extractLastNumber(line string) (int, error) {
	for firstPosition := len(line) - 1; firstPosition >= 0; firstPosition-- {
		for lastPosition := firstPosition + 1; lastPosition <= len(line); lastPosition++ {
			if number, err := tryConvertNumber(line[firstPosition:lastPosition]); err == nil {
				return number, nil
			}
		}
	}
	return 0, errors.New(line + " does not contain any numbers")
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

func (Day01) Part02(input string) string {
	lines := strings.Split(input, "\n")
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		first, err := extractFirstNumber(line)
		panicOnError(err)
		last, err := extractLastNumber(line)
		panicOnError(err)
		numbers = append(numbers, first*10+last)
	}

	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}

	return strconv.Itoa(sum)

}
