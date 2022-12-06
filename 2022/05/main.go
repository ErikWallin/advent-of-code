package main

import (
	"fmt"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) int {
	groups := common.ParseStringStringList(input, "\n\n", "\n")
	stacks := map[int][]rune{}

	for i := 1; i < 10; i++ {
		var si []rune

		for ri := len(groups[0]) - 2; ri >= 0; ri-- {
			row := groups[0][ri]
			c := rune(row[4*(i-1)+1])
			if c != ' ' {
				si = append(si, c)
			}
		}
		stacks[i] = si
	}

	for _, row := range groups[1] {
		instr := strings.Split(row, " ")
		amount := common.MustAtoi(instr[1])
		from := common.MustAtoi(instr[3])
		to := common.MustAtoi(instr[5])
		for i := 1; i <= amount; i++ {
			r := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], r)
		}
	}

	// Solution in text until run functions can return strings
	fmt.Printf("solution=")
	for i := 1; i < 10; i++ {
		fmt.Printf("%c", stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("\n")

	return 0
}

func run2(input string) int {
	groups := common.ParseStringStringList(input, "\n\n", "\n")
	stacks := map[int][]rune{}
	for i := 1; i < 10; i++ {
		var si []rune
		for ri := len(groups[0]) - 2; ri >= 0; ri-- {
			row := groups[0][ri]
			c := rune(row[4*(i-1)+1])
			if c != ' ' {
				si = append(si, c)
			}
		}
		stacks[i] = si
	}

	for _, row := range groups[1] {
		instr := strings.Split(row, " ")
		amount := common.MustAtoi(instr[1])
		from := common.MustAtoi(instr[3])
		to := common.MustAtoi(instr[5])
		r := stacks[from][len(stacks[from])-amount : len(stacks[from])]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		stacks[to] = append(stacks[to], r...)
	}

	// Solution in text until run functions can return strings
	fmt.Printf("solution=")
	for i := 1; i < 10; i++ {
		fmt.Printf("%c", stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("\n")

	return 0
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 5, submit, verbose)
}
