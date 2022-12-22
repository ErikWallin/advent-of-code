package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

type Monkey struct {
	name      string
	operation string
	first     string
	second    string
}

func yell(monkey Monkey) int {
	if monkey.operation == "" {
		return common.MustAtoi(monkey.first)
	} else if monkey.operation == "+" {
		return yell(monkeys[monkey.first]) + yell(monkeys[monkey.second])
	} else if monkey.operation == "-" {
		return yell(monkeys[monkey.first]) - yell(monkeys[monkey.second])
	} else if monkey.operation == "*" {
		return yell(monkeys[monkey.first]) * yell(monkeys[monkey.second])
	} else if monkey.operation == "/" {
		return yell(monkeys[monkey.first]) / yell(monkeys[monkey.second])
	}
	return 0
}

func containsHumn(monkey Monkey) bool {
	if monkey.name == "humn" {
		return true
	}
	if monkey.operation == "" {
		return false
	} else {
		return containsHumn(monkeys[monkey.first]) || containsHumn(monkeys[monkey.second])
	}
}

func getHumn(monkey Monkey, equals int) int {
	if monkey.name == "humn" {
		return equals
	}
	isLeft := containsHumn(monkeys[monkey.first])
	if isLeft {
		other := yell(monkeys[monkey.second])
		if monkey.operation == "+" {
			return getHumn(monkeys[monkey.first], equals-other)
		} else if monkey.operation == "-" {
			return getHumn(monkeys[monkey.first], equals+other)
		} else if monkey.operation == "*" {
			return getHumn(monkeys[monkey.first], equals/other)
		} else if monkey.operation == "/" {
			return getHumn(monkeys[monkey.first], equals*other)
		}
		return -1
	} else {
		other := yell(monkeys[monkey.first])
		if monkey.operation == "+" {
			return getHumn(monkeys[monkey.second], equals-other)
		} else if monkey.operation == "-" {
			return getHumn(monkeys[monkey.second], other-equals)
		} else if monkey.operation == "*" {
			return getHumn(monkeys[monkey.second], equals/other)
		} else if monkey.operation == "/" {
			return getHumn(monkeys[monkey.second], other/equals)
		}
		return -1
	}
}

var monkeys map[string]Monkey

func run1(input string) interface{} {
	monkeys = map[string]Monkey{}
	for _, m := range common.ParseStringList(input, "\n") {
		name := strings.Split(m, ": ")[0]
		parts := strings.Split((strings.Split(m, ": ")[1]), " ")
		if len(parts) == 1 {
			monkeys[name] = Monkey{name, "", parts[0], ""}
		} else {
			monkeys[name] = Monkey{name, parts[1], parts[0], parts[2]}
		}
	}
	return yell(monkeys["root"])
}

func run2(input string) interface{} {
	monkeys = map[string]Monkey{}
	for _, m := range common.ParseStringList(input, "\n") {
		name := strings.Split(m, ": ")[0]
		parts := strings.Split((strings.Split(m, ": ")[1]), " ")
		if len(parts) == 1 {
			first := parts[0]
			if name == "root" {
				first = "x"
			}
			monkeys[name] = Monkey{name, "", first, ""}
		} else {
			operation := parts[1]
			if name == "root" {
				operation = "="
			}
			monkeys[name] = Monkey{name, operation, parts[0], parts[2]}
		}
	}

	isLeft := containsHumn(monkeys[monkeys["root"].first])
	if isLeft {
		return getHumn(monkeys[monkeys["root"].first], yell(monkeys[monkeys["root"].second]))
	} else {
		return getHumn(monkeys[monkeys["root"].second], yell(monkeys[monkeys["root"].first]))
	}
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 21, submit)
}
