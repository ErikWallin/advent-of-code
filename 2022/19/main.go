package main

import (
	"fmt"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

type Blueprint struct {
	robots map[string]map[string]int
}

const ore = "ore"
const clay = "clay"
const obsidian = "obsidian"
const geode = "geode"

func parse(input string) []Blueprint {
	blueprints := []Blueprint{}
	for _, row := range common.ParseStringList(input, "\n") {
		row = strings.Split(row, ":")[1]
		row = strings.TrimSuffix(row, ".")
		robots := map[string]map[string]int{}
		var resource string
		for _, r := range strings.Split(row, ".") {
			r = strings.ReplaceAll(r, " Each ", "")
			r = strings.ReplaceAll(r, "robot costs ", "")
			r = strings.ReplaceAll(r, ".", "")
			r = strings.ReplaceAll(r, " and", "")
			parts := strings.Split(r, " ")
			resource = parts[0]
			robots[resource] = map[string]int{}
			for i := 1; i < len(parts)-1; i += 2 {
				robots[resource][parts[i+1]] = common.MustAtoi(parts[i])
			}
		}
		blueprints = append(blueprints, Blueprint{robots})
	}
	return blueprints
}

var currentRecord int
var cache map[string]map[string]int

func run1(input string) interface{} {
	blueprints := parse(input)

	sum := 0
	for i, b := range blueprints {
		currentRecord = 0
		cache = map[string]map[string]int{}
		money := map[string]int{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		robots := map[string]int{
			ore:      1,
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		sum += maximize(b, money, robots, 24) * (i + 1)
	}
	return sum
}

func mapClone(origin map[string]int) map[string]int {
	copy := make(map[string]int)
	for k, v := range origin {
		copy[k] = v
	}
	return copy
}

func run2(input string) interface{} {
	blueprints := parse(input)

	times := 3
	if common.ParseStringList(input, "\n")[0] == "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian." {
		times = 2
	}

	product := 1
	for _, b := range blueprints[:times] {
		currentRecord = 0
		cache = map[string]map[string]int{}
		money := map[string]int{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		robots := map[string]int{
			ore:      1,
			clay:     0,
			obsidian: 0,
			geode:    0,
		}
		geodes := maximize(b, money, robots, 32)
		product *= geodes
	}
	return product
}

func idFunc(robots map[string]int, t int) string {
	return fmt.Sprintf("%d,%d,%d,%d,%d", robots[ore], robots[clay], robots[obsidian], robots[geode], t)
}

func add(money map[string]int, robots map[string]int) map[string]int {
	moneyClone := mapClone(money)
	for k, v := range robots {
		moneyClone[k] += v
	}
	return moneyClone
}

func maximize(b Blueprint, money map[string]int, robots map[string]int, t int) int {
	id := idFunc(robots, t)
	m, ok := cache[id]
	if ok {
		if money[ore] <= m[ore] && money[clay] <= m[clay] && money[obsidian] <= m[obsidian] && money[geode] <= m[geode] {
			// Not better in any case
			return 0
		}
	}
	cache[id] = money

	// Optimizations
	canBuyGeode := money[ore] >= b.robots[geode][ore] && money[obsidian] >= b.robots[geode][obsidian]
	if t == 0 {
		result := money[geode]
		if result > currentRecord {
			currentRecord = result
		}
		return result
	} else if t == 1 {
		result := money[geode] + robots[geode]
		if result > currentRecord {
			currentRecord = result
		}
		return result
	} else if t == 2 && !canBuyGeode {
		result := money[geode] + robots[geode]*2
		if result > currentRecord {
			currentRecord = result
		}
		return result
	} else if t == 2 && canBuyGeode {
		result := money[geode] + robots[geode]*2 + 1
		if result > currentRecord {
			currentRecord = result
		}
		return result
	}

	// See if we can theoretically can reach highest record
	thisMax := money[geode] + robots[geode]*t
	if canBuyGeode {
		for ti := t; ti > 0; ti -= 1 {
			thisMax += ti - 1
		}
	} else {
		for ti := t - 1; ti > 0; ti -= 1 {
			thisMax += ti - 1
		}
	}
	if thisMax <= currentRecord {
		return 0
	}

	sums := []int{}
	if robots[obsidian] > 0 {
		robotsClone := mapClone(robots)
		robotsClone[geode]++
		moneyClone := mapClone(money)
		td := 1
		for {
			if moneyClone[ore] >= b.robots[geode][ore] && moneyClone[obsidian] >= b.robots[geode][obsidian] {
				moneyClone = add(moneyClone, robots)
				break
			}
			td++
			moneyClone = add(moneyClone, robots)
		}
		moneyClone[obsidian] -= b.robots[geode][obsidian]
		moneyClone[ore] -= b.robots[geode][ore]
		score := maximize(b, moneyClone, robotsClone, t-td)
		sums = append(sums, score)
	}
	// Buy obsidian
	if robots[clay] > 0 && robots[obsidian] < b.robots[geode][obsidian] {
		robotsClone := mapClone(robots)
		robotsClone[obsidian]++
		moneyClone := mapClone(money)
		td := 1
		for {
			if moneyClone[ore] >= b.robots[obsidian][ore] && moneyClone[clay] >= b.robots[obsidian][clay] {
				moneyClone = add(moneyClone, robots)
				break
			}
			td++
			moneyClone = add(moneyClone, robots)
		}
		moneyClone[clay] -= b.robots[obsidian][clay]
		moneyClone[ore] -= b.robots[obsidian][ore]
		score := maximize(b, moneyClone, robotsClone, t-td)
		sums = append(sums, score)
	}
	// Buy clay
	if robots[clay] < b.robots[obsidian][clay] {
		robotsClone := mapClone(robots)
		robotsClone[clay]++
		moneyClone := mapClone(money)
		td := 1
		for {
			if moneyClone[ore] >= b.robots[clay][ore] {
				moneyClone = add(moneyClone, robots)
				break
			}
			td++
			moneyClone = add(moneyClone, robots)
		}
		moneyClone[ore] -= b.robots[clay][ore]
		score := maximize(b, moneyClone, robotsClone, t-td)
		sums = append(sums, score)
	}
	// Buy ore
	if robots[ore] < common.MaxOfInts(b.robots[ore][ore], b.robots[clay][ore], b.robots[obsidian][ore], b.robots[geode][ore]) {
		robotsClone := mapClone(robots)
		robotsClone[ore]++
		moneyClone := mapClone(money)
		td := 1
		for {
			if moneyClone[ore] >= b.robots[ore][ore] {
				moneyClone = add(moneyClone, robots)
				break
			}
			td++
			moneyClone = add(moneyClone, robots)
		}
		moneyClone[ore] -= b.robots[ore][ore]
		score := maximize(b, moneyClone, robotsClone, t-td)
		sums = append(sums, score)
	}
	max := common.MaxOfInts(sums...)

	if max > currentRecord {
		currentRecord = max
	}
	return max
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 19, submit)
}
