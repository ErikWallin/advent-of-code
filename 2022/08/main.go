package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	matrix := common.ParseRuneListList(input)
	sum := len(matrix)*2 + len(matrix[0])*2 - 4
	for x := 1; x < len(matrix[0])-1; x++ {
		for y := 1; y < len(matrix)-1; y++ {
			seen := false
			seenLocal := true
			for xi := 0; xi < x; xi++ {
				if matrix[y][xi] >= matrix[y][x] {
					seenLocal = false
				}
			}
			if seenLocal {
				seen = true
			}

			seenLocal = true
			for xi := len(matrix[0]) - 1; xi > x; xi-- {
				if matrix[y][xi] >= matrix[y][x] {
					seenLocal = false
				}
			}
			if seenLocal {
				seen = true
			}

			seenLocal = true
			for yi := 0; yi < y; yi++ {
				if matrix[yi][x] >= matrix[y][x] {
					seenLocal = false
				}
			}
			if seenLocal {
				seen = true
			}

			seenLocal = true
			for yi := len(matrix[0]) - 1; yi > y; yi-- {
				if matrix[yi][x] >= matrix[y][x] {
					seenLocal = false
				}
			}
			if seenLocal {
				seen = true
			}
			if seen {
				sum++
			}
		}
	}
	return sum
}

func run2(input string) interface{} {
	matrix := common.ParseRuneListList(input)
	best := 0
	for x := 1; x < len(matrix[0])-1; x++ {
		for y := 1; y < len(matrix)-1; y++ {
			// left
			leftSum := 0
			for xi := x - 1; xi >= 0; xi-- {
				leftSum++
				if matrix[y][xi] >= matrix[y][x] {
					break
				}
			}
			// up
			upSum := 0
			for yi := y - 1; yi >= 0; yi-- {
				upSum++
				if matrix[yi][x] >= matrix[y][x] {
					break
				}
			}
			// right
			rightSum := 0
			for xi := x + 1; xi < len(matrix[0]); xi++ {
				rightSum++
				if matrix[y][xi] >= matrix[y][x] {
					break
				}
			}
			// down
			downSum := 0
			for yi := y + 1; yi < len(matrix); yi++ {
				downSum++
				if matrix[yi][x] >= matrix[y][x] {
					break
				}
			}
			scenic := leftSum * upSum * rightSum * downSum
			if scenic > best {
				best = scenic
			}
		}
	}
	return best
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 8, submit, verbose)
}
