package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

var tests = common.TestCases{
	{
		Input:         `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
		ExpectedPart1: 13,
		ExpectedPart2: 140,
	},
}
