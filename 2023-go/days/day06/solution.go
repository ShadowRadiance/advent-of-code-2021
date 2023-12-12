package day06

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	races := parseRaces(lines)

	waysToBeat := bruteForce(races)

	product := 0
	for _, num := range waysToBeat {
		if num > 0 {
			if product == 0 {
				product = num
			} else {
				product *= num
			}
		}
	}

	return strconv.Itoa(product)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	race := parseRaceWithoutSpaces(lines)

	waysToBeat := bruteForce([]Race{race})

	product := 0
	for _, num := range waysToBeat {
		if num > 0 {
			if product == 0 {
				product = num
			} else {
				product *= num
			}
		}
	}

	return strconv.Itoa(product)
}

type Race struct {
	time     int
	distance int
}

func (r Race) winHoldingFor(nSeconds int) bool {
	speed := nSeconds
	timeRemaining := r.time - nSeconds
	distance := timeRemaining * speed
	return distance > r.distance
}

func parseRaces(lines []string) []Race {
	_, timesStr, _ := strings.Cut(lines[0], ":")
	_, distancesStr, _ := strings.Cut(lines[1], ":")

	reSpaces := regexp.MustCompile(`\s+`)
	times := util.MapStringsToIntegers(
		reSpaces.Split(strings.TrimSpace(timesStr), -1))
	distances := util.MapStringsToIntegers(
		reSpaces.Split(strings.TrimSpace(distancesStr), -1))

	races := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		races[i].time = times[i]
		races[i].distance = distances[i]
	}

	return races
}

func parseRaceWithoutSpaces(lines []string) Race {
	_, timesStr, _ := strings.Cut(lines[0], ":")
	_, distancesStr, _ := strings.Cut(lines[1], ":")

	reSpaces := regexp.MustCompile(`\s+`)
	time := util.ConvertNumeric(
		strings.Join(
			reSpaces.Split(strings.TrimSpace(timesStr), -1),
			"",
		))
	distance := util.ConvertNumeric(
		strings.Join(
			reSpaces.Split(strings.TrimSpace(distancesStr), -1),
			"",
		))

	race := Race{time, distance}

	return race
}

func bruteForce(races []Race) []int {
	waysToBeat := make([]int, len(races))
	for iRace, race := range races {
		waysToBeatThisRace := 0
		for nSeconds := 0; nSeconds <= race.time; nSeconds++ {
			if race.winHoldingFor(nSeconds) {
				waysToBeatThisRace++
			}
		}
		waysToBeat[iRace] = waysToBeatThisRace
	}
	return waysToBeat
}

// we can improve on the brute force approach,
// but it works for both parts (test and run) so... maybe later

// for each race, find how many ways you could beat the distance within time
// [use a binary search to find the min_winner]
// [ways_to_win = (N+1) - (min_winner)*2]
// -- the distances follow a symmetric curve:
//   0: 0
//   1: a
//   2: aa
//   ...
//   n/2: aaaa…aaaa
//   ...
//   n-2: aa
//   n-1: a
//   n: 0
// )
