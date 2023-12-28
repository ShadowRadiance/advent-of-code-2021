package day19

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	workflows, remainingLines := parseWorkflows(lines)
	parts := parseParts(remainingLines)

	accepted := make([]Part, 0)
	rejected := make([]Part, 0)

	for _, part := range parts {
		result := part.runThroughWorkflows(workflows, "in")
		switch result {
		case "A":
			accepted = append(accepted, part)
		case "R":
			rejected = append(rejected, part)
		default:
			panic([]string{"WTF Megan!", result})
		}
	}

	ratings := util.Transform(accepted, func(item Part) int { return item.x + item.m + item.a + item.s })
	sum := util.Accumulate(ratings, func(total int, next int) int { return total + next })
	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

type Part struct {
	x, m, a, s int
}

type Rule = func(part Part) string
type Workflow = []Rule
type Workflows = map[string]Workflow

func parseWorkflows(lines []string) (Workflows, []string) {
	workflows := make(Workflows)

	last := 0
	for i, line := range lines {
		if len(line) == 0 {
			last = i
			break
		}
		matches := reWorkflow.FindStringSubmatch(line)
		name := matches[1]
		workflow := parseWorkflow(matches[2])
		workflows[name] = workflow
	}
	return workflows, lines[last+1:]
}

func parseWorkflow(line string) (workflow Workflow) {
	ruleStrings := strings.Split(line, ",")
	for _, ruleString := range ruleStrings {
		workflow = append(workflow, parseRule(ruleString))
	}
	return
}

var reComparison = regexp.MustCompile(`([xmas])(\D+)(\d+)`)
var reWorkflow = regexp.MustCompile(`(\w+)\{([^}]+)}`)

func parseRule(str string) Rule {
	chunks := strings.Split(str, ":")
	if len(chunks) == 1 {
		return func(part Part) string {
			return chunks[0]
		}
	}
	comparisonString := chunks[0]
	targetString := chunks[1]
	matches := reComparison.FindStringSubmatch(comparisonString)
	attribute := matches[1]
	comparator := matches[2]
	comparedValue := util.ConvertNumeric(matches[3])

	var reader func(Part) int
	switch attribute {
	case "x":
		reader = func(part Part) int { return part.x }
	case "m":
		reader = func(part Part) int { return part.m }
	case "a":
		reader = func(part Part) int { return part.a }
	case "s":
		reader = func(part Part) int { return part.s }
	}

	var comparison func(int) bool
	switch comparator {
	case "<":
		comparison = func(i int) bool { return i < comparedValue }
	case ">":
		comparison = func(i int) bool { return i > comparedValue }
	default:
		panic([]string{"DAMMIT WTF", comparator})
	}

	return func(part Part) string {
		if comparison(reader(part)) {
			return targetString
		}
		return ""
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
	line = line[1 : len(line)-1]
	chunks := strings.Split(line, ",")
	for _, chunk := range chunks {
		subChunks := strings.Split(chunk, "=")
		switch subChunks[0] {
		case "x":
			part.x = util.ConvertNumeric(subChunks[1])
		case "m":
			part.m = util.ConvertNumeric(subChunks[1])
		case "a":
			part.a = util.ConvertNumeric(subChunks[1])
		case "s":
			part.s = util.ConvertNumeric(subChunks[1])
		default:
			panic([]string{"DAMMIT WTF", chunk})
		}
	}
	return
}

func (part Part) runThroughWorkflows(flows Workflows, name string) (target string) {
	flow := flows[name]

	target = part.runThroughWorkflow(flow)
	switch target {
	case "A", "R":
		return
	default:
		return part.runThroughWorkflows(flows, target)
	}
}

func (part Part) runThroughWorkflow(flow Workflow) string {
	for _, rule := range flow {
		if result := rule(part); result != "" {
			return result
		}
	}
	panic("Processed all rules, but no target acquired")
}
