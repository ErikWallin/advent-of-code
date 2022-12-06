package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	score := 0
	for _, row := range rows {
		shapes := strings.Split(row, " ")
		if shapes[0] == "A" {
			if shapes[1] == "X" {
				score += 3
				score += 1
			} else if shapes[1] == "Y" {
				score += 6
				score += 2
			} else if shapes[1] == "Z" {
				score += 0
				score += 3
			}
		} else if shapes[0] == "B" {
			if shapes[1] == "X" {
				score += 0
				score += 1
			} else if shapes[1] == "Y" {
				score += 3
				score += 2
			} else if shapes[1] == "Z" {
				score += 6
				score += 3
			}
		} else if shapes[0] == "C" {
			if shapes[1] == "X" {
				score += 6
				score += 1
			} else if shapes[1] == "Y" {
				score += 0
				score += 2
			} else if shapes[1] == "Z" {
				score += 3
				score += 3
			}
		}
	}
	return score
}

func run2(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	score := 0
	for _, row := range rows {
		shapes := strings.Split(row, " ")
		if shapes[0] == "A" {
			if shapes[1] == "X" {
				score += 3
				score += 0
			} else if shapes[1] == "Y" {
				score += 1
				score += 3
			} else if shapes[1] == "Z" {
				score += 2
				score += 6
			}
		} else if shapes[0] == "B" {
			if shapes[1] == "X" {
				score += 1
				score += 0
			} else if shapes[1] == "Y" {
				score += 2
				score += 3
			} else if shapes[1] == "Z" {
				score += 3
				score += 6
			}
		} else if shapes[0] == "C" {
			if shapes[1] == "X" {
				score += 2
				score += 0
			} else if shapes[1] == "Y" {
				score += 3
				score += 3
			} else if shapes[1] == "Z" {
				score += 1
				score += 6
			}
		}
	}
	return score
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 2, submit, verbose)
}
