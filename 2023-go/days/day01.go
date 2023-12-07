package days

import (
	"strconv"
	"strings"
)

type Day01 struct{}

func (Day01) Part01(input string) string {
	lines := strings.Split(input, "\n")
	const numberChars = "1234567890"
	numbers := make([]int, 0, len(lines))

	for _, line := range lines {
		firstIdx := strings.IndexAny(line, numberChars)
		lastIdx := strings.LastIndexAny(line, numberChars)
		numberStr := string(line[firstIdx]) + string(line[lastIdx])
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}

	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}

	return strconv.Itoa(sum)
}

func (Day01) Part02(_ string) string { return "PENDING" }
