package main

import (
	"fmt"
	"math"

	"github.com/ErikWallin/advent-of-code/common"
)

type Blizzard struct {
	x         int
	y         int
	direction rune
}

type Coordinate struct {
	x int
	y int
}

var valley [][]rune

var startBlizzards []Blizzard
var blizzardsLoop map[int][]Blizzard
var blizzardsNoOfLoops int

func parse(input string) {
	valley = [][]rune{}
	startBlizzards = []Blizzard{}
	for y, r := range common.ParseRuneListList(input) {
		row := []rune{}
		for x, c := range r {
			if c == '#' {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
			if c != '#' && c != '.' {
				startBlizzards = append(startBlizzards, Blizzard{x, y, c})
			}
		}
		valley = append(valley, row)
	}
}

func initBlizzardsLoop() {
	blizzardsNoOfLoops = (len(valley) - 2) * (len(valley[0]) - 2)
	if len(valley) == len(valley[0]) {
		blizzardsNoOfLoops = len(valley) - 2
	}
	blizzardsLoop = map[int][]Blizzard{}
	blizzardsLoop[0] = startBlizzards
	for i := 1; i < blizzardsNoOfLoops; i++ {
		loop := []Blizzard{}
		for _, pb := range blizzardsLoop[i-1] {
			b := pb
			if b.direction == '>' {
				b.x++
				if b.x == len(valley[0])-1 {
					b.x = 1
				}
			} else if b.direction == '<' {
				b.x--
				if b.x == 0 {
					b.x = len(valley[0]) - 2
				}
			}
			if b.direction == 'v' {
				b.y++
				if b.y == len(valley)-1 {
					b.y = 1
				}
			} else if b.direction == '^' {
				b.y--
				if b.y == 0 {
					b.y = len(valley) - 2
				}
			}
			loop = append(loop, b)
		}
		blizzardsLoop[i] = loop
	}
}

var record int
var cache map[string]int

func run1(input string) interface{} {
	parse(input)
	initBlizzardsLoop()
	start := Coordinate{1, 0}
	end := Coordinate{len(valley[0]) - 2, len(valley) - 1}
	record = math.MaxInt
	cache = map[string]int{}
	reachGoal(start, end, start, 0)
	return record
}

func reachGoal(start Coordinate, end Coordinate, expedition Coordinate, t int) {
	manhattanDistanceToEnd := int(math.Abs(float64(end.x-expedition.x)) + math.Abs(float64(end.y-expedition.y)))
	if manhattanDistanceToEnd <= 1 {
		if t+1 < record {
			record = t + 1
		}
		return
	}
	cacheId := fmt.Sprintf("%d%d%d", expedition.x, expedition.y, t%blizzardsNoOfLoops)
	val, ok := cache[cacheId]
	if ok && t >= val {
		return
	}
	cache[cacheId] = t
	if t+manhattanDistanceToEnd >= record {
		return
	}
	moves := []Coordinate{{1, 0}, {0, 1}, {0, 0}, {0, -1}, {-1, 0}}
	if start.x > end.x {
		moves = []Coordinate{{-1, 0}, {0, -1}, {0, 0}, {0, 1}, {1, 0}}
	}
	for _, m := range moves {
		okToMove := true
		next := Coordinate{expedition.x + m.x, expedition.y + m.y}
		if !(expedition == start) && (next.x <= 0 || next.x >= len(valley[0])-1 || next.y <= 0 || next.y >= len(valley)-1) {
			okToMove = false
		}
		if okToMove {
			for _, b := range blizzardsLoop[(t+1)%blizzardsNoOfLoops] {
				if next.x == b.x && next.y == b.y {
					okToMove = false
					break
				}
			}
		}
		if okToMove {
			reachGoal(start, end, next, t+1)
		}
	}
}

func run2(input string) interface{} {
	parse(input)
	initBlizzardsLoop()
	start := Coordinate{1, 0}
	end := Coordinate{len(valley[0]) - 2, len(valley) - 1}
	record = math.MaxInt
	cache = map[string]int{}
	reachGoal(start, end, start, 0)
	time := record
	record = math.MaxInt
	cache = map[string]int{}
	reachGoal(end, start, end, time)
	time = record
	record = math.MaxInt
	cache = map[string]int{}
	reachGoal(start, end, start, time)
	return record
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 24, submit)
}
