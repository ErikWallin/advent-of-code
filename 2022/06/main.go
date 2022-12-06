package main

import (
	"github.com/ErikWallin/advent-of-code/common"
	mapset "github.com/deckarep/golang-set/v2"
)

func run1(input string) int {
	for i := 0; i < len(input)-5; i++ {
		word := input[i : i+4]
		set := mapset.NewSet[rune]()
		for _, c := range word {
			set.Add(c)
		}
		if set.Cardinality() == 4 {
			return i + 4
		}
	}
	return 0
}

func run2(input string) int {
	for i := 0; i < len(input)-15; i++ {
		word := input[i : i+14]
		set := mapset.NewSet[rune]()
		for _, c := range word {
			set.Add(c)
		}
		if set.Cardinality() == 14 {
			return i + 14
		}
	}
	return 0
}

func main() {
	submit := true
	verbose := false
	common.Run(run1, run2, tests, 2022, 6, submit, verbose)
}
