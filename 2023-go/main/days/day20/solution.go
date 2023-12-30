package day20

import (
	"fmt"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Solution struct{}

type Pulse struct {
	from   string
	target string
	level  string
}

func (p Pulse) destructure() (string, string, string) { return p.from, p.target, p.level }

type Module struct {
	name       string
	kind       string
	outputs    []string
	flipState  string
	conjMemory map[string]string
}

func (m *Module) serialize() string {
	memory := ""
	if m.kind == "%" {
		memory = m.flipState
	} else if m.kind == "&" {
		for name, state := range m.conjMemory {
			memory += fmt.Sprintf("[%s=%s]", name, state)
		}
	}
	return fmt.Sprintf("{%s:%s[%s]:%s}", m.kind, m.name, strings.Join(m.outputs, ","), memory)
}

func (m *Module) receive(pulse Pulse) (pulses []Pulse) {
	var outgoing string
	if m.kind == "%" {
		if pulse.level == "hi" {
			return
		}
		if m.flipState == "on" {
			m.flipState = "off"
			outgoing = "lo"
		} else {
			m.flipState = "on"
			outgoing = "hi"
		}
	} else { // m.kind == "&"
		m.conjMemory[pulse.from] = pulse.level
		outgoing = "lo"
		for _, level := range m.conjMemory {
			if level == "lo" {
				outgoing = "hi"
				break
			}
		}
	}
	for _, output := range m.outputs {
		pulses = append(pulses, Pulse{m.name, output, outgoing})
	}
	return
}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	counts := map[string]int{"lo": 0, "hi": 0}
	modules, broadcastTargets := parseModules(lines)

	for i := 0; i < 1000; i++ {
		counts["lo"]++ // the initial button press to the broadcaster
		pulses := aq.New()
		for _, target := range broadcastTargets {
			pulses.Enqueue(Pulse{from: "broadcaster", target: target, level: "lo"})
		}
		for !pulses.Empty() {
			p, _ := pulses.Dequeue()
			pulse := p.(Pulse)
			counts[pulse.level]++
			if module, ok := modules[pulse.target]; ok {
				newPulses := module.receive(pulse)
				for _, newPulse := range newPulses {
					pulses.Enqueue(newPulse)
				}
			}
		}
	}

	return strconv.Itoa(counts["hi"] * counts["lo"])
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	modules, broadcastTargets := parseModules(lines)

	// TRUTH: There is only one "rx", we want to know when it gets a lo signal
	buttonPresses := 0
	//rxModule := modules["rx"]
	// ASSUMPTION 1: There is only one thing that feeds into "rx" and it is a conjunction (&)
	var feederModules []string
	for name, module := range modules {
		for _, target := range module.outputs {
			if target == "rx" {
				feederModules = append(feederModules, name)
			}
		}
	}
	if len(feederModules) != 1 {
		panic("Assumption 1 failed: more (or fewer) than one module feeds into rx")
	}
	feederModule := feederModules[0]
	if modules[feederModule].kind != "&" {
		panic("Assumption 1 failed: feeder module is not a conjunction")
	}

	// ASSUMPTION 2: Since that is a & (conjunction), i.e. it only outputs lo when all inputs are hi,
	//               the things that send to the & will each send hi on independent cycles,
	//               thus the LCM of those cycles is how many presses it will take.

	seenSendHiToFeeder := map[string]int{}
	for name, module := range modules {
		for _, target := range module.outputs {
			if target == feederModule {
				seenSendHiToFeeder[name] = 0
			}
		}
	}

	cycleLengthBySender := map[string]int{}
	for {
		buttonPresses++
		pulses := aq.New()
		for _, target := range broadcastTargets {
			pulses.Enqueue(Pulse{from: "broadcaster", target: target, level: "lo"})
		}
		for !pulses.Empty() {
			p, _ := pulses.Dequeue()
			pulse := p.(Pulse)
			if module, ok := modules[pulse.target]; ok {
				if module.name == feederModule && pulse.level == "hi" {
					seenSendHiToFeeder[pulse.from]++
					if _, ok := cycleLengthBySender[pulse.from]; !ok {
						cycleLengthBySender[pulse.from] = buttonPresses
					} else {
						if buttonPresses != seenSendHiToFeeder[pulse.from]*cycleLengthBySender[pulse.from] {
							panic("Assumption 2 failed: sender not on a cycle")
						}
					}
					if util.All(util.MapValues(seenSendHiToFeeder), func(value int) bool { return value > 0 }) {
						lcm := 1
						for _, cycleLength := range cycleLengthBySender {
							lcm = util.LowestCommonMultiple(lcm, cycleLength)
						}
						return strconv.Itoa(lcm)
					}
				}

				newPulses := module.receive(pulse)
				for _, newPulse := range newPulses {
					pulses.Enqueue(newPulse)
				}
			}
		}
	}
}

var (
	reLine = regexp.MustCompile(`([%&]?)(\w+) -> (.*)`)
)

func parseModules(lines []string) (modules map[string]*Module, broadcastTargets []string) {
	modules = map[string]*Module{}

	for _, line := range lines {
		matches := reLine.FindStringSubmatch(line)
		typeName, name, outputsStr := matches[1], matches[2], matches[3]
		outputs := strings.Split(outputsStr, ", ")
		if name == "broadcaster" {
			broadcastTargets = outputs
		} else {
			modules[name] = &Module{
				name:       name,
				kind:       typeName,
				outputs:    outputs,
				flipState:  "off",
				conjMemory: map[string]string{},
			}
		}
	}

	// initialize the prevStates on conjunctions
	for name, module := range modules {
		for _, output := range module.outputs {
			if targetModule, ok := modules[output]; ok {
				if targetModule.kind == "&" {
					targetModule.conjMemory[name] = "lo"
				}
			}
		}
	}

	return
}
