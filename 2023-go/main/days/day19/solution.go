package day19

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	workflows := parseWorkflows(strings.Split(sections[0], "\n"))
	parts := parseParts(strings.Split(sections[1], "\n"))

	sum := 0
	for _, part := range parts {
		if accept(part, workflows, "in") {
			sum += part["x"] + part["m"] + part["a"] + part["s"]
		}
	}
	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	workflows := parseWorkflows(strings.Split(sections[0], "\n"))

	intervals := map[string]util.Interval{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	total := count(workflows, intervals, "in")
	return strconv.Itoa(total)
}

// vvvvvvvvvvvvvvvvvvvvvvvv PART 2 vvvvvvvvvvvvvvvvvvvvvvvv //

type Interval = util.Interval
type Intervals = map[string]util.Interval

func count(workflows Workflows, intervals Intervals, name string) int {
	if name == "R" {
		return 0
	}
	if name == "A" {
		product := 1
		for _, v := range intervals {
			product *= v.Length()
		}
		return product
	}

	total := 0

	workflow := workflows[name]
	brokeOut := false
	for _, rule := range workflow.rules {
		var tHalf, fHalf Interval
		// split the interval into 2:
		if rule.comparison == "<" {
			tHalf = Interval{Start: intervals[rule.attribute].Start, Final: rule.compareValue - 1}
			fHalf = Interval{Start: rule.compareValue, Final: intervals[rule.attribute].Final}
		} else {
			fHalf = Interval{Start: intervals[rule.attribute].Start, Final: rule.compareValue}
			tHalf = Interval{Start: rule.compareValue + 1, Final: intervals[rule.attribute].Final}
		}
		if !tHalf.Invalid() {
			newIntervals := copyIntervals(intervals)
			newIntervals[rule.attribute] = tHalf
			total += count(workflows, newIntervals, rule.target)
		}
		if !fHalf.Invalid() {
			newIntervals := copyIntervals(intervals)
			newIntervals[rule.attribute] = fHalf
			intervals = newIntervals // set up for next rule check
		} else {
			// false half was empty
			brokeOut = true
			break
		}
	}
	if !brokeOut {
		// ran out of rules - use fallback
		total += count(workflows, intervals, workflow.fallback)
	}

	return total
}

func copyIntervals(intervals Intervals) Intervals {
	newIntervals := Intervals{}
	for key, interval := range intervals {
		newIntervals[key] = interval
	}
	return newIntervals
}

// ^^^^^^^^^^^^^^^^^^^^^^^^ PART 2 ^^^^^^^^^^^^^^^^^^^^^^^^ //

type Part = map[string]int

var comps = map[string]func(int, int) bool{
	"<": func(a, b int) bool { return a < b },
	">": func(a, b int) bool { return a > b },
}

type Rule struct {
	attribute    string
	comparison   string
	compareValue int
	target       string
}
type Workflow struct {
	rules    []Rule
	fallback string
}
type Workflows = map[string]Workflow

func parseWorkflows(lines []string) (workflows Workflows) {
	workflows = make(Workflows)

	for _, line := range lines {
		matches := reWorkflow.FindStringSubmatch(line)
		workflows[matches[1]] = parseWorkflow(matches[2])
	}
	return workflows
}

func parseWorkflow(line string) (workflow Workflow) {
	ruleStrings := strings.Split(line, ",")
	workflow = Workflow{fallback: ruleStrings[len(ruleStrings)-1], rules: make([]Rule, 0)}
	for _, ruleString := range ruleStrings[:len(ruleStrings)-1] {
		workflow.rules = append(workflow.rules, parseRule(ruleString))
	}
	return
}

var reComparison = regexp.MustCompile(`([xmas])(\D+)(\d+)`)
var reWorkflow = regexp.MustCompile(`(\w+)\{([^}]+)}`)

func parseRule(str string) Rule {
	chunks := strings.Split(str, ":")
	comparisonString := chunks[0]
	targetString := chunks[1]
	matches := reComparison.FindStringSubmatch(comparisonString)

	return Rule{
		attribute:    matches[1],
		comparison:   matches[2],
		compareValue: util.ConvertNumeric(matches[3]),
		target:       targetString,
	}
}

func parseParts(lines []string) (parts []Part) {
	for _, line := range lines {
		parts = append(parts, parsePart(line))
	}
	return
}

func parsePart(line string) (part Part) {
	// strip off the {}
	pairs := strings.Split(line[1:len(line)-1], ",")
	part = make(Part)
	for _, pair := range pairs {
		chunks := strings.Split(pair, "=")
		part[chunks[0]] = util.ConvertNumeric(chunks[1])
	}
	return
}

func accept(part Part, flows Workflows, name string) bool {
	if name == "R" {
		return false
	}
	if name == "A" {
		return true
	}

	flow := flows[name]
	for _, rule := range flow.rules {
		if comps[rule.comparison](part[rule.attribute], rule.compareValue) {
			return accept(part, flows, rule.target)
		}
	}
	return accept(part, flows, flow.fallback)
}
