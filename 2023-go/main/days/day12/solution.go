package day12

import (
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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	// input: "???.### 1,1,3"
	// record: {
	//  list: "???.###????.###????.###????.###????.###"
	//  counts: 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3
	// }
	records := unfoldRecords(parseRecords(lines))
	// fmt.Printf("%+v\n", records[0])

	sum := 0
	for _, record := range records {
		sum += validPermutations(record)
	}

	// Elapsed: 6.883779042s
	return strconv.Itoa(sum)
}

type Record struct {
	list   string
	counts []int
}

func parseRecords(lines []string) (records []Record) {
	for _, line := range lines {
		records = append(records, parseRecord(line))
	}
	return
}

func parseRecord(line string) (record Record) {
	ss := strings.Split(line, " ")
	record.list = ss[0]
	record.counts = util.MapStringsToIntegers(strings.Split(ss[1], ","))
	return
}

func unfoldRecords(records []Record) []Record {
	for i, record := range records {
		records[i] = unfoldRecord(record)
	}
	return records
}

func unfoldRecord(record Record) Record {
	repeatedSprings := strings.Repeat(record.list+"?", 5)
	repeatedSprings = repeatedSprings[0 : len(repeatedSprings)-1]

	size := len(record.counts)
	repeatedCounts := make([]int, len(record.counts)*5)
	for i := 0; i < 5; i++ {
		copy(
			repeatedCounts[i*size:(i+1)*size],
			record.counts,
		)
	}

	return Record{
		list:   repeatedSprings,
		counts: repeatedCounts,
	}
}

func validPermutations(record Record) int {
	// BLATANTLY STOLEN FROM https://github.com/tmo1/adventofcode/blob/main/2023/12.py
	// AND TRANSLATED FROM PYTHON TO GO
	// Fast and... I need to split this up to understand it better
	ways := 0
	positions := map[int]int{0: 1}
	for i, contiguous := range record.counts {
		newPositions := map[int]int{}
		for k, v := range positions {
			futureCounts := record.counts[i+1:]
			futureCountSum := 0
			if len(futureCounts) > 0 {
				futureCountSum = util.Accumulate(futureCounts, func(tot, ct int) int { return tot + ct })
			}
			for n := k; n < len(record.list)-futureCountSum+len(futureCounts); n++ {
				if n+contiguous-1 < len(record.list) && !strings.Contains(record.list[n:n+contiguous], ".") {
					if (i == len(record.counts)-1 && !strings.Contains(record.list[n+contiguous:], "#")) || (i < len(record.counts)-1 && n+contiguous < len(record.list) && rune(record.list[n+contiguous]) != '#') {
						existing, ok := newPositions[n+contiguous+1]
						if ok {
							newPositions[n+contiguous+1] = existing + v
						} else {
							newPositions[n+contiguous+1] = v
						}
					}
				}
				if rune(record.list[n]) == '#' {
					break
				}
			}
		}
		positions = newPositions
	}
	for _, v := range positions {
		ways += v
	}
	return ways

}
