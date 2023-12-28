package day05

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Interval = util.Interval

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	seeds, remainingLines := parseSeeds(lines)
	seedToSoilMap, remainingLines := parseMap(remainingLines)
	soilToFertilizerMap, remainingLines := parseMap(remainingLines)
	fertilizerToWaterMap, remainingLines := parseMap(remainingLines)
	waterToLightMap, remainingLines := parseMap(remainingLines)
	lightToTemperatureMap, remainingLines := parseMap(remainingLines)
	temperatureToHumidityMap, remainingLines := parseMap(remainingLines)
	humidityToLocationMap, remainingLines := parseMap(remainingLines)

	seedsToLocations := make(map[int]int)
	smallestLocation := math.MaxInt

	for _, seed := range seeds {

		soil := seedToSoilMap.read(seed)
		fertilizer := soilToFertilizerMap.read(soil)
		water := fertilizerToWaterMap.read(fertilizer)
		light := waterToLightMap.read(water)
		temperature := lightToTemperatureMap.read(light)
		humidity := temperatureToHumidityMap.read(temperature)
		location := humidityToLocationMap.read(humidity)

		seedsToLocations[seed] = location
		if location < smallestLocation {
			smallestLocation = location
		}
	}

	return strconv.Itoa(smallestLocation)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	seedIntervals, remainingLines := parseSeedIntervals(lines)

	seedToSoilMap, remainingLines := parseMap(remainingLines)
	soilToFertilizerMap, remainingLines := parseMap(remainingLines)
	fertilizerToWaterMap, remainingLines := parseMap(remainingLines)
	waterToLightMap, remainingLines := parseMap(remainingLines)
	lightToTemperatureMap, remainingLines := parseMap(remainingLines)
	temperatureToHumidityMap, remainingLines := parseMap(remainingLines)
	humidityToLocationMap, remainingLines := parseMap(remainingLines)

	maps := []Map{
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	}
	for _, m := range maps {
		sort.Sort(ByMappingSourceStart(m.mappings))
	}

	ranges := make(map[Interval]Interval)
	for _, seedInterval := range seedIntervals {
		ranges[seedInterval] = seedInterval
	}

	for _, m := range maps {
		applyMapToRanges(m, ranges)
	}

	smallestLocation := math.MaxInt
	for _, iLocation := range ranges {
		if iLocation.Start < smallestLocation {
			smallestLocation = iLocation.Start
		}
	}

	return strconv.Itoa(smallestLocation)
}

func applyMapToRanges(translator Map, ranges map[Interval]Interval) {

	var keys []Interval
	for key := range ranges {
		keys = append(keys, key)
	}

	var keysToDelete []Interval
	var subMaps []map[Interval]Interval

	for seedInterval, targetInterval := range ranges {
		keysToDelete = append(keysToDelete, seedInterval)
		subMaps = append(subMaps, splitIntervals(seedInterval, targetInterval, translator))
	}

	for _, interval := range keysToDelete {
		delete(ranges, interval)
	}

	for _, subMap := range subMaps {
		for k, v := range subMap {
			ranges[k] = v
		}
	}
}

func transformRange(interval Interval, m Map) (newInterval Interval) {
	// precondition: sub-range must fit within single element of Map.mappings
	for _, mapping := range m.mappings {
		if interval.Start >= mapping.src && interval.Final <= mapping.src+mapping.len-1 {
			change := mapping.dst - mapping.src
			newInterval.Start = interval.Start + change
			newInterval.Final = interval.Final + change
			return
		}
	}
	newInterval = interval
	return
	// panic("precondition not met: sub-range must fit within single element of Map.mappings")
}

func splitIntervals(seedInterval Interval, targetInterval Interval, translator Map) map[Interval]Interval {
	result := make(map[Interval]Interval)

	intervalMappings := make([]map[Interval]Interval, 0)
	for _, mapping := range translator.mappings {
		intervalMappings = append(intervalMappings, mapping.toIntervalMapping()) // { (56,92):(60,96), (93,96):(56,59)
	}

	var intervalMappingsKeys []Interval
	for _, mapping := range intervalMappings {
		for iKey := range mapping {
			intervalMappingsKeys = append(intervalMappingsKeys, iKey) // [ (56,92), (93,96) ]
		}
	}
	sort.Sort(ByStart(intervalMappingsKeys))

	// targetInterval is before all mappings
	// tiS-tiF < kS-kF...
	if targetInterval.Final < intervalMappingsKeys[0].Start {
		result[seedInterval] = targetInterval
		return result
	}

	// targetInterval is after all mappings
	// ...kS-kF < tiS-tiF
	if targetInterval.Start > intervalMappingsKeys[len(intervalMappingsKeys)-1].Final {
		result[seedInterval] = targetInterval
		return result
	}

	// target interval is completely inside a mapping
	// ...kS < tiS-tiF < kF...
	for _, key := range intervalMappingsKeys {
		if targetInterval.Start >= key.Start && targetInterval.Final <= key.Final {
			result[seedInterval] = transformRange(targetInterval, translator)
			return result
		}
	}

	target := Interval{Start: targetInterval.Start, Final: targetInterval.Final}
	seed := Interval{Start: seedInterval.Start, Final: seedInterval.Final}
	for _, key := range intervalMappingsKeys {
		// kS-KF < [tS-tF]
		if key.Final < target.Start {
			continue // target not affected by this key
		}
		// [tS-tF] < kS-kF
		if key.Start > target.Final {
			continue // target not affected by this key
		}
		// tS < kS <= tF <= kF
		// tS < kS <= kF <= tF
		// break off "tS...kS-1" as a new interval (leaving [ts=ks...tf]<=kf) or [ts=ks...kf<tf]
		if target.Start < key.Start {
			newTarget := Interval{Start: target.Start, Final: key.Start - 1}
			newSeed := Interval{Start: seed.Start, Final: seed.Start + newTarget.Length() - 1}
			result[newSeed] = transformRange(newTarget, translator)
			// remove used portion of interval & seed
			target.Start += newTarget.Length()
			seed.Start += newTarget.Length()
		}
		// [tS=kS <= tF] <= kF
		// [tS=kS <= kF < tF]
		if target.Start == key.Start {
			if target.Final <= key.Final {
				// [tS=kS <= tF] <= kF
				// done with target "tS...tF"
				result[seed] = transformRange(target, translator)
				return result
			} else { // key.Final < target.Final
				// [tS=kS <= kF < tF]
				// break off "tS...kF" as a new interval (leaving ts...tf which might cross another key)
				newTarget := Interval{Start: target.Start, Final: key.Final}
				newSeed := Interval{Start: seed.Start, Final: seed.Start + newTarget.Length() - 1}
				result[newSeed] = transformRange(newTarget, translator)
				// remove used portion of interval & seed
				target.Start += newTarget.Length()
				seed.Start += newTarget.Length()
				continue // done with this key
			}
		}
		// kS < tS < kF < tF
		if target.Start > key.Start {
			if target.Final <= key.Final {
				// [kS < tS <= tF] <= kF
				// done with target "tS...tF"
				result[seed] = transformRange(target, translator)
				return result
			} else {
				// kS < [tS <= kF < tF]
				// break off "tS...kF" as a new interval (leaving ts...tf which might cross another key)
				newTarget := Interval{Start: target.Start, Final: key.Final}
				newSeed := Interval{Start: seed.Start, Final: seed.Start + newTarget.Length() - 1}
				result[newSeed] = transformRange(newTarget, translator)
				// remove used portion of interval & seed
				target.Start += newTarget.Length()
				seed.Start += newTarget.Length()
				continue // done with this key
			}
		}
	}
	// if there is any target left, add it
	if !target.Invalid() {
		result[seed] = target
	}

	return result
}

func parseSeeds(lines []string) (seeds []int, remainingLines []string) {
	seeds = util.MapStringsToIntegers(strings.Split(lines[0][7:], " "))
	remainingLines = lines[2:]
	return
}

type Mapping struct {
	dst int
	src int
	len int
}

func (m Mapping) srcEnd() int {
	return m.src + m.len - 1
}

func (m Mapping) dstEnd() int {
	return m.dst + m.len - 1
}

func (m Mapping) toIntervalMapping() map[Interval]Interval {
	return map[Interval]Interval{
		{Start: m.src, Final: m.srcEnd()}: {Start: m.dst, Final: m.dstEnd()},
	}
}

type ByMappingSourceStart []Mapping

func (a ByMappingSourceStart) Len() int           { return len(a) }
func (a ByMappingSourceStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMappingSourceStart) Less(i, j int) bool { return a[i].src < a[j].src }

type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

type Map struct {
	name     string
	mappings []Mapping
}

func (m Map) read(in int) int {
	for _, mapping := range m.mappings {
		if in >= mapping.src && in <= mapping.src+mapping.len {
			return mapping.dst + (in - mapping.src)
		}
	}
	return in
}

func parseMap(lines []string) (Map, []string) {
	map_ := Map{}
	map_.name = lines[0][:len(lines[0])-5]
	remainingLines := lines[1:]
	for i, line := range remainingLines {
		if len(line) == 0 {
			return map_, remainingLines[i+1:]
		}
		map_.mappings = append(map_.mappings, parseMapping(line))
	}
	return map_, remainingLines[0:0]
}

func parseMapping(line string) Mapping {
	// drs srs len
	integers := util.MapStringsToIntegers(strings.Split(line, " "))
	return Mapping{integers[0], integers[1], integers[2]}
}

func parseSeedIntervals(lines []string) (seedIntervals []Interval, remainingLines []string) {
	numbers := util.MapStringsToIntegers(strings.Split(lines[0][7:], " "))
	chunks := util.ChunkIntSlice(numbers, 2)

	for _, chunk := range chunks {
		seedIntervals = append(seedIntervals, Interval{
			Start: chunk[0],
			Final: chunk[0] + chunk[1] - 1,
		})
	}

	remainingLines = lines[2:]
	return
}

// seeds: 79-92 / 55-67 (seeds)
//
// MAP:      SRC-(SRC+LEN-1):DST-(DST+LEN-1) / ETC...               => MAP[SEED RANGE]TARGET RANGE
//
// seeds:                                                           => 79-92:79-92                                         / 55-67:55-67                                                         (seeds:seeds)
//
// soils:    50-97:52-99 / 98-99:50-51                              => 79-92:t(79-92)                                      / 55-67:t(55-67)                                                      (seeds:soils)
// soils:    50-97:52-99 / 98-99:50-51                              => 79-92:81-94                                         / 55-67:57-69                                                         (seeds:soils)
//
// ferts:    15-51: 0-36 / 52-53:37-38 /  0-14:39-53                => 79-92:t(81-94)                                      / 55-67:t(57-69)                                                      (seeds:ferts)
// ferts:    15-51: 0-36 / 52-53:37-38 /  0-14:39-53                => 79-92:81-94                                         / 55-67:57-69                                                         (seeds:ferts)
//
// water:    53-60:49-56 / 11-52: 0-41 /  0- 6:42:48 /  7-10:57-60  => 79-92:t(81-94)                                      / 55-67:t(57-69)                                                      (seeds:water)
// water:    53-60:49-56 / 11-52: 0-41 /  0- 6:42:48 /  7-10:57-60  => 79-92:t(81-94)                                      / [55-58:t(57-60) / 59-67:t(61-69)]                                   (seeds:water)
// water:    53-60:49-56 / 11-52: 0-41 /  0- 6:42:48 /  7-10:57-60  => 79-92:81-94                                         / [55-58:53-56    / 59-67:61-69   ]                                   (seeds:water)
//
// light:    18-24:88-94 / 25-94:18-87                              => 79-92:t(81-94)                                      / 55-58:t(53-56)  / 59-67:t(61-69)                                    (seeds:light)
// light:    18-24:88-94 / 25-94:18-87                              => 79-92:74-87                                         / 55-58:46-49     / 59-67:54-62                                       (seeds:light)
//
// temps:    77-99:45-77 / 45-63:81-99 / 64-76:68-80                => 79-92:t(74-87)                                      / 55-58:t(46-49)  / 59-67:t(54-62)                                    (seeds:temps)
// temps:    77-99:45-77 / 45-63:81-99 / 64-76:68-80                => [79-81:t(74-76) / 82-92:t(77-87)]                   / 55-58:t(46-49)  / 59-67:t(54-62)                                    (seeds:temps)
// temps:    77-99:45-77 / 45-63:81-99 / 64-76:68-80                => [79-81:78-80    / 82-92:45-55   ]                   / 55-58:82-85     / 59-67:90-98                                       (seeds:temps)
//
// humid:    69-70: 0- 1 / 0-68:1-69                                => 79-81:t(78-80)  / 82-92:t(45-55)                    / 55-58:t(82-85)  / 59-67:t(90-98)                                    (seeds:humid)
// humid:    69-70: 0- 1 / 0-68:1-69                                => 79-81:78-80     / 82-92:46-56                       / 55-58:82-85     / 59-67:90-98                                       (seeds:humid)
//
// locns:    56-92:60-96 / 93-96:56-59                              => 79-81:t(78-80)  / 82-92:t(46-56)                    / 55-58:t(82-85)  / 59-67:t(90-98)                                     (seeds:locns)
// locns:    56-92:60-96 / 93-96:56-59                              => 79-81:t(78-80)  / [82-91:t(46-55) / 92-92:t(56-56)] / 55-58:t(82-85)  / [59-61:t(90-92) / 62-65:t(93-96) / 66-67:t(97-98)] (seeds:locns)
// locns:    56-92:60-96 / 93-96:56-59                              => 79-81:82-84     / [82-91:46-55    / 92-92:60-60]    / 55-58:86-89     / [59-61:94-96    / 62-65:56-59    / 66-67:97-98   ] (seeds:locns)
