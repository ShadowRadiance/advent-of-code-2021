package day20

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/emirpasic/gods/queues"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type HiLo int

const (
	Lo HiLo = iota
	Hi
)

type Pulse struct {
	from  string
	to    string
	level HiLo
}

func (p Pulse) destructure() (string, string, HiLo) { return p.from, p.to, p.level }
func dequeue(q queues.Queue) (Pulse, bool) {
	p, ok := q.Dequeue()
	return p.(Pulse), ok
}

type FlipFlop struct {
	name    string
	state   HiLo
	outputs []string
}
type Conjunction struct {
	// remember to initialize prevStates for ALL inputs
	name       string
	prevStates map[string]HiLo
	outputs    []string
}
type Broadcast struct {
	name    string
	outputs []string
}
type Dummy struct {
	name string
}

func (m FlipFlop) receive(pulse Pulse) (pulses []Pulse) {
	if pulse.level == Lo {
		if m.state == Hi {
			m.state = Lo
		} else {
			m.state = Hi
		}
		for _, output := range m.outputs {
			pulses = append(pulses, Pulse{from: m.name, to: output, level: m.state})
		}
	}
	return
}
func (m Conjunction) receive(pulse Pulse) (pulses []Pulse) {
	from, _, level := pulse.destructure()
	m.prevStates[from] = level

	outputLevel := Lo
	for _, v := range m.prevStates {
		if v == Lo {
			outputLevel = Hi
			break
		}
	}
	for _, output := range m.outputs {
		pulses = append(pulses, Pulse{from: m.name, to: output, level: outputLevel})
	}
	return
}
func (m Broadcast) receive(pulse Pulse) (pulses []Pulse) {
	for _, output := range m.outputs {
		pulses = append(pulses, Pulse{from: m.name, to: output, level: pulse.level})
	}
	return
}
func (m Dummy) receive(_ Pulse) (pulses []Pulse) {
	return
}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	modules := parseModules(lines)
	pulses := aq.New()
	pulses.Enqueue(Pulse{from: "button", to: "broadcaster", level: Lo})

	count := 0
	//for pulse, ok := dequeue(pulses); ok; count++ {
	//	from, to, level := pulse.destructure()
	//	toModule := modules[to]
	//
	//}

	spew.Dump(modules)

	return strconv.Itoa(count)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

var (
	reLine = regexp.MustCompile(`([%&]?)(\w+) -> (.*)`)
)

func parseModules(lines []string) map[string]interface{} {
	modules := map[string]interface{}{}
	outputNames := map[string]bool{}

	for _, line := range lines {
		matches := reLine.FindStringSubmatch(line)
		typeName, name, outputsStr := matches[1], matches[2], matches[3]
		outputs := strings.Split(outputsStr, ", ")
		for _, output := range outputs {
			outputNames[output] = true
		}

		switch typeName {
		case "%":
			modules[name] = FlipFlop{
				name:    name,
				state:   Lo,
				outputs: outputs,
			}
		case "&":
			modules[name] = Conjunction{
				name:       name,
				prevStates: make(map[string]HiLo),
				outputs:    outputs,
			}
		default:
			modules[name] = Broadcast{
				name:    name,
				outputs: outputs,
			}
		}
	}

	// there may be outputs that were Dummies
	for name := range outputNames {
		if _, ok := modules[name]; !ok {
			modules[name] = Dummy{name: name}
		}
	}

	// initialize the inputs on conjunctions
	for _, m := range modules {
		if b, ok := m.(Broadcast); ok {
			for _, output := range b.outputs {
				if conj, ok := modules[output].(Conjunction); ok {
					conj.prevStates[b.name] = Lo
				}
			}
		}
		if f, ok := m.(FlipFlop); ok {
			for _, output := range f.outputs {
				if conj, ok := modules[output].(Conjunction); ok {
					conj.prevStates[f.name] = Lo
				}
			}
		}
		if c, ok := m.(Conjunction); ok {
			for _, output := range c.outputs {
				if conj, ok := modules[output].(Conjunction); ok {
					conj.prevStates[c.name] = Lo
				}
			}
		}
	}

	return modules
}
