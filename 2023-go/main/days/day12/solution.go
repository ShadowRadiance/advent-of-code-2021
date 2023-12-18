package day12

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	records := parseRecords(lines)

	sum := 0
	for _, record := range records {
		sum += validPermutations(record)
	}
	// fmt.Printf("%+v\n", records)

	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

type Record struct {
	SpringList string
	Counts     []int
}

func parseRecords(lines []string) (records []Record) {
	for _, line := range lines {
		records = append(records, parseRecord(line))
	}
	return
}

func parseRecord(line string) (record Record) {
	ss := strings.Split(line, " ")
	record.SpringList = ss[0]
	record.Counts = util.MapStringsToIntegers(strings.Split(ss[1], ","))
	return
}

func validPermutations(record Record) int {
	return recursive(record.SpringList, record.Counts, 0)
}

func recursive(list string, counts []int, currentIndex int) int {
	if len(list) == currentIndex {
		if isValid(list, counts) {
			return 1
		} else {
			return 0
		}
	}

	if list[currentIndex] == '?' {
		return recursive(
			list[:currentIndex]+"#"+list[currentIndex+1:],
			counts,
			currentIndex+1,
		) + recursive(
			list[:currentIndex]+"."+list[currentIndex+1:],
			counts,
			currentIndex+1,
		)
	} else {
		return recursive(list, counts, currentIndex+1)
	}
}

func isValid(list string, counts []int) bool {
	re := regexp.MustCompile(`#+`)
	blocks := re.FindAllString(string(list), -1)
	// list #.#.### => blocks [ #, #, ### ]
	if len(blocks) != len(counts) {
		return false
	}
	blockCounts := util.Transform(blocks, func(item string) int { return len(item) })

	for i, v := range blockCounts {
		if v != counts[i] {
			return false
		}
	}
	return true
}
