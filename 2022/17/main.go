package main

import (
	"fmt"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	moves := []rune(input)
	rocks := [][][]rune{
		{[]rune("####")},
		{[]rune(".#."), []rune("###"), []rune(".#.")},
		{[]rune("###"), []rune("..#"), []rune("..#")},
		{[]rune("#"), []rune("#"), []rune("#"), []rune("#")},
		{[]rune("##"), []rune("##")},
	}
	chamber, _ := simulate([][]rune{}, 0, 2022, moves, rocks)
	return len(chamber)
}

func collision(chamber [][]rune, rock [][]rune, x, y int) bool {
	for yr, row := range rock {
		for xr, rr := range row {
			posX := x + xr
			posY := y + yr
			if rr == '#' && posY < len(chamber) && chamber[posY][posX] == '#' {
				return true
			}
		}
	}
	return false
}

func addToChamber(chamber [][]rune, rock [][]rune, x, y int) [][]rune {
	for yr, row := range rock {
		posY := yr + y
		if posY >= len(chamber) {
			chamber = append(chamber, []rune("......."))
		}
		for xr, rr := range row {
			posX := xr + x
			if rr == '#' {
				chamber[posY][posX] = rr
			}
		}
	}
	return chamber
}

func printChamber(chamber [][]rune) {
	for y := len(chamber) - 1; y >= 0; y-- {
		fmt.Printf("|%s|\n", string(chamber[y]))
	}
	println("+-------+")
}

func printChamberTop(chamber [][]rune) {
	for y := len(chamber) - 1; y >= len(chamber)-10; y-- {
		fmt.Printf("|%s|\n", string(chamber[y]))
	}
	println("|ooooooo|")
}

func run2(input string) interface{} {
	moves := []rune(input)
	rocks := [][][]rune{
		{[]rune("####")},
		{[]rune(".#."), []rune("###"), []rune(".#.")},
		{[]rune("###"), []rune("..#"), []rune("..#")},
		{[]rune("#"), []rune("#"), []rune("#"), []rune("#")},
		{[]rune("##"), []rune("##")},
	}

	noOfShapesInLoop := 0
	start := 1000
	chamber, m := simulate([][]rune{}, 0, start, moves, rocks)
	c1000 := chamber
	m1000 := m
	for {
		chamber, m = simulate(chamber, m, 5, moves, rocks)
		noOfShapesInLoop += 5
		if m%len(moves) == m1000%len(moves) {
			break
		}
		if noOfShapesInLoop > 10000 {
			println("Could not find loop")
			break
		}
	}
	lengthLoop := len(chamber) - len(c1000)
	leftover := (1000000000000 - start) % noOfShapesInLoop
	cFinish, _ := simulate(chamber, m, leftover, moves, rocks)
	length := len(c1000) + ((1000000000000-start)/noOfShapesInLoop)*lengthLoop + len(cFinish) - len(chamber)
	return length
}

func simulate(chamber [][]rune, moveIndex int, times int, moves []rune, rocks [][][]rune) ([][]rune, int) {
	m := moveIndex
	for i := 0; i < times; i++ {
		rock := rocks[i%5]
		x := 2
		y := len(chamber) + 3
		for {
			move := moves[m%len(moves)]
			m++
			xn := x + 1
			if move == '<' {
				xn = x - 1
			}
			if xn < 0 || xn+len(rock[0]) > 7 || collision(chamber, rock, xn, y) {
				xn = x
			}
			x = xn
			yn := y - 1
			if yn < 0 || collision(chamber, rock, x, yn) {
				chamber = addToChamber(chamber, rock, x, y)
				break
			}
			y = yn
		}
	}
	return chamber, m
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 17, submit)
}
