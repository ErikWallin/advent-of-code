package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type Pipe struct {
	valve    string
	flowRate int
	tunnels  []string
}

var digCache map[string]int

func digCacheId1(valve string, time int, opened mapset.Set[string]) string {
	openedSlice := opened.ToSlice()
	sort.Slice(openedSlice, func(i, j int) bool {
		return openedSlice[i] < openedSlice[j]
	})
	return fmt.Sprintf("%s%d%v", valve, time, openedSlice)
}

func digCacheId2(valve1 string, valve2 string, time int, opened mapset.Set[string]) string {
	openedSlice := opened.ToSlice()
	sort.Slice(openedSlice, func(i, j int) bool {
		return openedSlice[i] < openedSlice[j]
	})
	valves := []string{valve1, valve2}
	sort.Slice(valves, func(i, j int) bool {
		return valves[i] < valves[j]
	})
	return fmt.Sprintf("%s%d%v", valves, time, openedSlice)
}

func dig1(current Pipe, opened mapset.Set[string], time int) int {
	if time == 0 {
		return 0
	}
	time--
	sums := []int{}
	if !opened.Contains(current.valve) && current.flowRate != 0 {
		openedClone := opened.Clone()
		openedClone.Add(current.valve)
		addition := current.flowRate * time
		cacheId := digCacheId1(current.valve, time, openedClone)
		cacheHit, ok := digCache[cacheId]
		if ok {
			sums = append(sums, addition+cacheHit)
		} else {
			val := dig1(current, openedClone, time)
			digCache[cacheId] = val
			sums = append(sums, addition+val)
		}
	}
	for _, nextPipe := range current.tunnels {
		cacheId := digCacheId1(pipes[nextPipe].valve, time, opened.Clone())
		cacheHit, ok := digCache[cacheId]
		if ok {
			sums = append(sums, cacheHit)
		} else {
			val := dig1(pipes[nextPipe], opened.Clone(), time)
			digCache[cacheId] = val
			sums = append(sums, val)
		}
	}
	return common.MaxOfInts(sums...)
}

func run1(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	pipes = map[string]Pipe{}
	for _, row := range list {
		parts := strings.Split(row, " to valve")
		tunnels := strings.Split(strings.TrimPrefix(strings.ReplaceAll(parts[1], "s", ""), " "), ", ")
		pipe := Pipe{
			strings.ReplaceAll(strings.Split(row, " has flow rate")[0], "Valve ", ""),
			common.MustAtoi(strings.Split(strings.Split(row, ";")[0], "=")[1]),
			tunnels}
		pipes[pipe.valve] = pipe
	}

	digCache = map[string]int{}
	res := dig1(pipes["AA"], mapset.NewSet[string](), 30)
	//fmt.Printf("res=%v\n", res)
	return res
}

func dig2(current1 Pipe, current2 Pipe, opened mapset.Set[string], time int, maxOpened int, currentSum int) int {
	if time == 0 {
		return 0
	}
	time--
	sums := []int{}
	if opened.Cardinality() == maxOpened {
		// Cannot get more points
		return 0
	}
	if currentSum < 100 && time < 22 {
		return 0
	}
	if currentSum < 500 && time < 18 {
		return 0
	}
	if currentSum < 1000 && time < 14 {
		return 0
	}
	if currentSum < 1500 && time < 8 {
		return 0
	}
	for _, c1 := range append(current1.tunnels, current1.valve) {
		for _, c2 := range append(current2.tunnels, current2.valve) {
			if (c1 == current1.valve && (opened.Contains(c1) || current1.flowRate == 0)) ||
				(c2 == current2.valve && (opened.Contains(c2) || current2.flowRate == 0)) {
				// Do not open again
				// Do not open valves that give not result
				continue
			}
			if c1 == current1.valve && c1 == c2 {
				// Do not open same valve
				continue
			}
			if c1 == current1.valve && c2 == current2.valve {
				openedClone := opened.Clone()
				openedClone.Add(c1)
				openedClone.Add(c2)
				addition := (current1.flowRate + current2.flowRate) * time
				cacheId := digCacheId2(c1, c2, time, openedClone)
				cacheHit, ok := digCache[cacheId]
				if ok {
					sums = append(sums, addition+cacheHit)
				} else {
					val := dig2(current1, current2, openedClone, time, maxOpened, currentSum+addition)
					digCache[cacheId] = val
					sums = append(sums, addition+val)
				}
			} else if c1 == current1.valve && c2 != current2.valve {
				openedClone := opened.Clone()
				openedClone.Add(c1)
				addition := current1.flowRate * time
				cacheId := digCacheId2(c1, pipes[c2].valve, time, openedClone)
				cacheHit, ok := digCache[cacheId]
				if ok {
					sums = append(sums, addition+cacheHit)
				} else {
					val := dig2(current1, pipes[c2], openedClone, time, maxOpened, currentSum+addition)
					digCache[cacheId] = val
					sums = append(sums, addition+val)
				}
			} else if c1 != current1.valve && c2 == current2.valve {
				openedClone := opened.Clone()
				openedClone.Add(c2)
				addition := current2.flowRate * time
				cacheId := digCacheId2(pipes[c1].valve, c2, time, openedClone)
				cacheHit, ok := digCache[cacheId]
				if ok {
					sums = append(sums, addition+cacheHit)
				} else {
					val := dig2(pipes[c1], current2, openedClone, time, maxOpened, currentSum+addition)
					digCache[cacheId] = val
					sums = append(sums, addition+val)
				}
			} else if c1 != current1.valve && c2 != current2.valve {
				openedClone := opened.Clone()
				cacheId := digCacheId2(pipes[c1].valve, pipes[c2].valve, time, openedClone)
				cacheHit, ok := digCache[cacheId]
				if ok {
					sums = append(sums, cacheHit)
				} else {
					val := dig2(pipes[c1], pipes[c2], openedClone, time, maxOpened, currentSum)
					digCache[cacheId] = val
					sums = append(sums, val)
				}
			}
		}
	}
	return common.MaxOfInts(sums...)
}

func run2(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	pipes = map[string]Pipe{}
	maxOpened := 0
	for _, row := range list {
		parts := strings.Split(row, " to valve")
		tunnels := strings.Split(strings.TrimPrefix(strings.ReplaceAll(parts[1], "s", ""), " "), ", ")
		flowRate := common.MustAtoi(strings.Split(strings.Split(row, ";")[0], "=")[1])
		pipe := Pipe{
			strings.ReplaceAll(strings.Split(row, " has flow rate")[0], "Valve ", ""),
			flowRate,
			tunnels}
		pipes[pipe.valve] = pipe
		if flowRate > 0 {
			maxOpened++
		}
	}

	digCache = map[string]int{}
	res := dig2(pipes["AA"], pipes["AA"], mapset.NewSet[string](), 26, maxOpened, 0)
	return res
}

var pipes map[string]Pipe

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 16, submit, verbose)
}
