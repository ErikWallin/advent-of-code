package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

var tests = common.TestCases{
	{
		Input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
		ExpectedPart1: 95437,
		ExpectedPart2: 24933642,
	},
}
