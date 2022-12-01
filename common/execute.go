package common

import (
	"fmt"
	"time"
)

func Run(run1 func(string) int, run2 func(string) int, tests TestCases, year int, day int, submit bool, verbose bool) (int, int) {
	if tests != nil {
		tests.Run(run1, run2, !verbose)
	}
	input := ReadFile(year, day, "puzzle.txt")
	start := time.Now()
	part1 := run1(input)
	part2 := run2(input)
	elapsed := time.Since(start)
	fmt.Printf("PART1: %v\nPART2: %v\n", part1, part2)
	fmt.Printf("Program took %s\n", elapsed)
	if submit && part2 != 0 {
		Submit(year, day, 2, part2)
	} else if submit && part1 != 0 {
		Submit(year, day, 1, part1)
	}
	return part1, part2
}
