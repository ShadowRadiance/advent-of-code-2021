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

	// Elapsed: 6.883779042s
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
	return recursiveCount(record)
}

func recursiveCount(record Record) int {
	if !strings.Contains(record.SpringList, "?") {
		return util.BoolInt(record.isValid())
	} else {
		return recursiveCount(Record{SpringList: strings.Replace(record.SpringList, "?", ".", 1), Counts: record.Counts}) +
			recursiveCount(Record{SpringList: strings.Replace(record.SpringList, "?", "#", 1), Counts: record.Counts})
	}
}

func (record Record) isValid() bool {
	re := regexp.MustCompile(`#+`)
	blocks := re.FindAllString(record.SpringList, -1)
	// list #.#.### => blocks [ #, #, ### ]
	if len(blocks) != len(record.Counts) {
		return false
	}
	blockCounts := util.Transform(blocks, func(item string) int { return len(item) })

	for i, v := range blockCounts {
		if v != record.Counts[i] {
			return false
		}
	}
	return true
}
