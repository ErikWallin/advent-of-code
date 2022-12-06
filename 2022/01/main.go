package main

import (
	"sort"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	elves := common.ParseStringList(input, "\n\n")
	heaviest := 0
	for _, elf := range elves {
		cals := common.ParseIntList(elf, "\n")
		sum := 0
		for _, cal := range cals {
			sum = sum + cal
		}
		if sum > heaviest {
			heaviest = sum
		}
	}
	return heaviest
}

func run2(input string) interface{} {
	elves := common.ParseStringList(input, "\n\n")
	var ans []int
	for _, elf := range elves {
		cals := common.ParseIntList(elf, "\n")
		sum := 0
		for _, cal := range cals {
			sum = sum + cal
		}
		ans = append(ans, sum)
	}
	sort.Ints(ans)
	ans = common.ReverseSlice(ans)
	return ans[0] + ans[1] + ans[2]
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 1, submit, verbose)
}
