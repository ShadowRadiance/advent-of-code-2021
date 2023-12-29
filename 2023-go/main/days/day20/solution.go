package day20

import (
	"strings"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
	"regexp"
)

type Solution struct{}

type Module interface {
	Name() string
	Inputs() []string
	Outputs() []string
	SetInputs([]string)
}

type ModuleData struct {
	name    string
	inputs  []string
	outputs []string
}

type FlipFlop struct {
	ModuleData
	state bool
}

type Conjunction struct {
	ModuleData
	lastInputLevels []bool
}

type Broadcaster struct {
	ModuleData
}

type Output struct {
	ModuleData
}

func (data *ModuleData) Name() string              { return data.name }
func (data *ModuleData) Inputs() []string          { return data.inputs }
func (data *ModuleData) Outputs() []string         { return data.outputs }
func (data *ModuleData) SetInputs(inputs []string) { (*data).inputs = inputs }

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	var modules = parseModules(lines)
	spew.Dump(modules)

	return "PENDING"
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

func parseModules(lines []string) map[string]Module {
	modules := map[string]Module{}

	for _, line := range lines {
		module := parseModule(line)
		modules[module.Name()] = module
	}

	for _, module := range modules {
		for _, output := range module.Outputs() {
      if _, ok := modules[output]; !ok {
        modules[output] = &Output{ModuleData{name: output, outputs: make([]string,0), inputs: make([]string, 0)}}
      }
    }
  }

	for name, module := range modules {
		for _, output := range module.Outputs() {
			modules[output].SetInputs(append(modules[output].Inputs(), name))
		}
	}
	return modules
}

func parseModule(line string) Module {
	var module Module

	re := regexp.MustCompile(`([%&]?)(\w+) -> (.*)`)
	matches := re.FindStringSubmatch(line)

	typeName, name, outputsStr := matches[1], matches[2], matches[3]
	outputs := strings.Split(outputsStr, ", ")
	moduleData := ModuleData{
		name:    name,
		outputs: outputs,
    inputs: make([]string, 0),
	}

	switch typeName {
	case "%":
		module = &FlipFlop{moduleData, false}
	case "&":
		module = &Conjunction{moduleData, make([]bool, 0)}
	default:
		module = &Broadcaster{moduleData}
	}

	return module
}
