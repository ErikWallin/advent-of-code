package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

var tests = common.TestCases{
	{
		Input:         `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`,
		ExpectedPart1: 152,
		ExpectedPart2: 301,
	},
}
