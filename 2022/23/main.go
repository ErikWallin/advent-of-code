package main

import (
	"github.com/ErikWallin/advent-of-code/common"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Coordinate struct {
	x int
	y int
}

var elves map[int]Coordinate
var directions map[int][]rune

func run1(input string) interface{} {
	elves = map[int]Coordinate{}
	directions = map[int][]rune{}
	index := 0
	for y, row := range common.ParseRuneListList(input) {
		for x, c := range row {
			if c == '#' {
				elves[index] = Coordinate{x, y}
				directions[index] = []rune{'N', 'S', 'W', 'E'}
				index++
			}
		}
	}

	for round := 1; round <= 10; round++ {
		proposals := map[int]Coordinate{}
		for ie, elf := range elves {
			alone := true
			for x := elf.x - 1; x <= elf.x+1; x++ {
				for y := elf.y - 1; y <= elf.y+1; y++ {
					if !(elf.x == x && elf.y == y) && slices.Contains(maps.Values(elves), Coordinate{x, y}) {
						alone = false
					}
				}
			}
			if !alone {
				contains := false
				for _, d := range directions[ie] {
					if d == 'N' {
						contains = false
						for _, oe := range elves {
							if oe.y == elf.y-1 && oe.x >= elf.x-1 && oe.x <= elf.x+1 {
								contains = true
								break
							}
						}
						if !contains {
							proposals[ie] = Coordinate{elf.x, elf.y - 1}
							break
						}
					}
					if d == 'S' {
						contains = false
						for _, oe := range elves {
							if oe.y == elf.y+1 && oe.x >= elf.x-1 && oe.x <= elf.x+1 {
								contains = true
								break
							}
						}
						if !contains {
							proposals[ie] = Coordinate{elf.x, elf.y + 1}
							break
						}
					}
					if d == 'W' {
						contains = false
						for _, oe := range elves {
							if oe.x == elf.x-1 && oe.y >= elf.y-1 && oe.y <= elf.y+1 {
								contains = true
								break
							}
						}
						if !contains {
							proposals[ie] = Coordinate{elf.x - 1, elf.y}
							break
						}
					}
					if d == 'E' {
						contains = false
						for _, oe := range elves {
							if oe.x == elf.x+1 && oe.y >= elf.y-1 && oe.y <= elf.y+1 {
								contains = true
								break
							}
						}
						if !contains {
							proposals[ie] = Coordinate{elf.x + 1, elf.y}
							break
						}
					}
				}
				if contains {
					proposals[ie] = Coordinate{elf.x, elf.y}
				}
			} else {
				proposals[ie] = Coordinate{elf.x, elf.y}
			}
		}

		newElves := map[int]Coordinate{}
		for ip, p := range proposals {
			collision := false
			for ipo, po := range proposals {
				if ip != ipo && p == po {
					newElves[ip] = elves[ip]
					collision = true
				}
			}
			if !collision {
				newElves[ip] = proposals[ip]
			}
			firstDirection := directions[ip][0]
			copy(directions[ip][:], directions[ip][1:])
			directions[ip][3] = firstDirection
		}
		elves = newElves
	}
	min, max := minMax(maps.Values(elves))
	return (max.x-min.x+1)*(max.y-min.y+1) - len(maps.Values(elves))
}

func minMax(elves []Coordinate) (Coordinate, Coordinate) {
	min := Coordinate{1000, 1000}
	max := Coordinate{-1, -1}

	for _, elf := range elves {
		if elf.x < min.x {
			min.x = elf.x
		}
		if elf.x > max.x {
			max.x = elf.x
		}
		if elf.y < min.y {
			min.y = elf.y
		}
		if elf.y > max.y {
			max.y = elf.y
		}
	}
	return min, max
}

func run2(input string) interface{} {
	elves = map[int]Coordinate{}
	directions = map[int][]rune{}
	index := 0
	for y, row := range common.ParseRuneListList(input) {
		for x, c := range row {
			if c == '#' {
				elves[index] = Coordinate{x, y}
				directions[index] = []rune{'N', 'S', 'W', 'E'}
				index++
			}
		}
	}

	round := 1
	for {
		moved := false
		proposals := map[int]Coordinate{}
		for ie, elf := range elves {
			alone := true
			for x := elf.x - 1; x <= elf.x+1; x++ {
				for y := elf.y - 1; y <= elf.y+1; y++ {
					if !(elf.x == x && elf.y == y) && slices.Contains(maps.Values(elves), Coordinate{x, y}) {
						alone = false
					}
				}
			}
			if !alone {
				contains := false
				for _, d := range directions[ie] {
					if d == 'N' {
						contains = false
						for _, oe := range elves {
							if oe.y == elf.y-1 && oe.x >= elf.x-1 && oe.x <= elf.x+1 {
								contains = true
								break
							}
						}
						if !contains {
							moved = true
							proposals[ie] = Coordinate{elf.x, elf.y - 1}
							break
						}
					}
					if d == 'S' {
						contains = false
						for _, oe := range elves {
							if oe.y == elf.y+1 && oe.x >= elf.x-1 && oe.x <= elf.x+1 {
								contains = true
								break
							}
						}
						if !contains {
							moved = true
							proposals[ie] = Coordinate{elf.x, elf.y + 1}
							break
						}
					}
					if d == 'W' {
						contains = false
						for _, oe := range elves {
							if oe.x == elf.x-1 && oe.y >= elf.y-1 && oe.y <= elf.y+1 {
								contains = true
								break
							}
						}
						if !contains {
							moved = true
							proposals[ie] = Coordinate{elf.x - 1, elf.y}
							break
						}
					}
					if d == 'E' {
						contains = false
						for _, oe := range elves {
							if oe.x == elf.x+1 && oe.y >= elf.y-1 && oe.y <= elf.y+1 {
								contains = true
								break
							}
						}
						if !contains {
							moved = true
							proposals[ie] = Coordinate{elf.x + 1, elf.y}
							break
						}
					}
				}
				if contains {
					proposals[ie] = Coordinate{elf.x, elf.y}
				}
			} else {
				proposals[ie] = Coordinate{elf.x, elf.y}
			}
		}

		newElves := map[int]Coordinate{}
		for ip, p := range proposals {
			collision := false
			for ipo, po := range proposals {
				if ip != ipo && p == po {
					newElves[ip] = elves[ip]
					collision = true
				}
			}
			if !collision {
				newElves[ip] = proposals[ip]
			}
			firstDirection := directions[ip][0]
			copy(directions[ip][:], directions[ip][1:])
			directions[ip][3] = firstDirection
		}
		elves = newElves
		if !moved {
			return round
		}
		round++
	}
}

func printMap() {
	min, max := minMax(maps.Values(elves))
	for y := min.y - 1; y <= max.y+1; y++ {
		for x := min.x - 1; x <= max.x+1; x++ {
			contains := false
			for _, e := range elves {
				if e.x == x && e.y == y {
					contains = true
					break
				}
			}
			if contains {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func main() {
	submit := true
	common.Run(run1, run2, tests, 2022, 23, submit)
}
