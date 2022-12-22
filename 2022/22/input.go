package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

var tests = common.TestCases{
	{
		Input:         `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`,
		ExpectedPart1: 6032,
		ExpectedPart2: 5031,
	},
}
