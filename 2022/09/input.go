package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

var tests = common.TestCases{
	{
		Input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
		ExpectedPart1: 13,
		ExpectedPart2: 1,
	},
	{
		Input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
		ExpectedPart1: 0,
		ExpectedPart2: 36,
	},
}
