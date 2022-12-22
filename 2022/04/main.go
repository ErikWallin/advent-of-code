package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	sum := 0
	for _, row := range list {
		elves := strings.Split(row, ",")
		elf1 := strings.Split(elves[0], "-")
		elf2 := strings.Split(elves[1], "-")
		if common.MustAtoi(elf1[0]) >= common.MustAtoi(elf2[0]) && common.MustAtoi(elf1[1]) <= common.MustAtoi(elf2[1]) {
			sum++
		} else if common.MustAtoi(elf1[0]) <= common.MustAtoi(elf2[0]) && common.MustAtoi(elf1[1]) >= common.MustAtoi(elf2[1]) {
			sum++
		}
	}
	return sum
}

func run2(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	sum := 0
	for _, row := range list {
		elves := strings.Split(row, ",")
		elf1 := strings.Split(elves[0], "-")
		elf2 := strings.Split(elves[1], "-")
		if common.MustAtoi(elf1[0]) > common.MustAtoi(elf2[1]) || common.MustAtoi(elf1[1]) < common.MustAtoi(elf2[0]) {
			sum++
		}
	}
	return len(list) - sum
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 4, submit)
}
