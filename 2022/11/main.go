package main

import (
	"sort"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

type Monkey struct {
	Items         []int
	Op            string
	OpArg1        string
	OpArg2        string
	TestDivisible int
	OnTrue        int
	OnFalse       int
	Inspections   int
}

func run1(input string) interface{} {
	groups := common.ParseStringStringList(input, "\n\n", "\n")
	monkeys := []*Monkey{}
	for _, monkey := range groups {
		itemStr := strings.Split(strings.Split(monkey[1], ": ")[1], ", ")
		items := []int{}
		for _, s := range itemStr {
			items = append(items, common.MustAtoi(s))
		}
		opStr := strings.Split(strings.Split(monkey[2], "= ")[1], " ")
		op := opStr[1]
		opArg1 := opStr[0]
		opArg2 := opStr[2]
		testDivisible := common.MustAtoi(strings.Split(monkey[3], "by ")[1])
		onTrue := common.MustAtoi(strings.Split(monkey[4], "monkey ")[1])
		onFalse := common.MustAtoi(strings.Split(monkey[5], "monkey ")[1])
		monkeys = append(monkeys, &Monkey{items, op, opArg1, opArg2, testDivisible, onTrue, onFalse, 0})
	}

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.Inspections++
				var opArg1 int
				if monkey.OpArg1 == "old" {
					opArg1 = item
				} else {
					opArg1 = common.MustAtoi(monkey.OpArg1)
				}
				var opArg2 int
				if monkey.OpArg2 == "old" {
					opArg2 = item
				} else {
					opArg2 = common.MustAtoi(monkey.OpArg2)
				}
				var worryLevel int
				if monkey.Op == "*" {
					worryLevel = opArg1 * opArg2
				} else if monkey.Op == "+" {
					worryLevel = opArg1 + opArg2
				}
				worryLevel = worryLevel / 3
				if worryLevel%monkey.TestDivisible == 0 {
					monkeys[monkey.OnTrue].Items = append(monkeys[monkey.OnTrue].Items, worryLevel)
				} else {
					monkeys[monkey.OnFalse].Items = append(monkeys[monkey.OnFalse].Items, worryLevel)
				}
			}
			monkey.Items = []int{}
		}
	}
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Ints(inspections)
	inspections = common.ReverseSlice(inspections)
	return inspections[0] * inspections[1]
}

func run2(input string) interface{} {
	groups := common.ParseStringStringList(input, "\n\n", "\n")
	monkeys := []*Monkey{}
	for _, monkey := range groups {
		itemStr := strings.Split(strings.Split(monkey[1], ": ")[1], ", ")
		items := []int{}
		for _, s := range itemStr {
			items = append(items, common.MustAtoi(s))
		}
		opStr := strings.Split(strings.Split(monkey[2], "= ")[1], " ")
		op := opStr[1]
		opArg1 := opStr[0]
		opArg2 := opStr[2]
		testDivisible := common.MustAtoi(strings.Split(monkey[3], "by ")[1])
		onTrue := common.MustAtoi(strings.Split(monkey[4], "monkey ")[1])
		onFalse := common.MustAtoi(strings.Split(monkey[5], "monkey ")[1])
		monkeys = append(monkeys, &Monkey{items, op, opArg1, opArg2, testDivisible, onTrue, onFalse, 0})
	}

	divider := 1
	for _, monkey := range monkeys {
		divider *= monkey.TestDivisible
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.Inspections++
				var opArg1 int
				if monkey.OpArg1 == "old" {
					opArg1 = item
				} else {
					opArg1 = common.MustAtoi(monkey.OpArg1)
				}
				var opArg2 int
				if monkey.OpArg2 == "old" {
					opArg2 = item
				} else {
					opArg2 = common.MustAtoi(monkey.OpArg2)
				}
				var worryLevel int
				if monkey.Op == "*" {
					worryLevel = opArg1 * opArg2
				} else if monkey.Op == "+" {
					worryLevel = opArg1 + opArg2
				}
				worryLevel = worryLevel % divider
				if worryLevel%monkey.TestDivisible == 0 {
					monkeys[monkey.OnTrue].Items = append(monkeys[monkey.OnTrue].Items, worryLevel)
				} else {
					monkeys[monkey.OnFalse].Items = append(monkeys[monkey.OnFalse].Items, worryLevel)
				}
			}
			monkey.Items = []int{}
		}
	}
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Ints(inspections)
	inspections = common.ReverseSlice(inspections)
	return inspections[0] * inspections[1]
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 11, submit, verbose)
}
