package common

import (
	"fmt"
	"os"
)

type TestCases []TestCase

type TestCase struct {
	Input         string
	ExpectedPart1 int
	ExpectedPart2 int
}

func (t TestCases) Run(fn1 func(string) interface{}, fn2 func(string) interface{}) {
	for _, test := range t {
		part1I := fn1(test.Input)
		passedPart1 := part1I == test.ExpectedPart1 || test.ExpectedPart1 == 0
		if passedPart1 && test.ExpectedPart1 != 0 {
			fmt.Println(" - PART1: ", part1I, " correct")
		}
		if !passedPart1 {
			fmt.Println(" - PART1: ", part1I, " but expected ", test.ExpectedPart1)
			os.Exit(1)
		}

		part2I := fn2(test.Input)
		passedPart2 := part2I == test.ExpectedPart2 || test.ExpectedPart2 == 0
		if passedPart2 && test.ExpectedPart2 != 0 {
			fmt.Println(" - PART2: ", part2I, " correct")
		}
		if !passedPart2 {
			fmt.Println(" - PART2: ", part2I, " but expected ", test.ExpectedPart2)
			os.Exit(1)
		}
	}
}
