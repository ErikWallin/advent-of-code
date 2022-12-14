package common

import (
	"strings"
)

// Parse input to slice of integers, i.e. []int
// Works for both csv (sep = ",") and row delimited (sep = "\n")
func ParseIntList(input string, sep string) []int {
	lines := strings.Split(input, sep)
	list := make([]int, len(lines))
	for i, line := range lines {
		list[i] = MustAtoi(line)
	}
	return list
}

// Parse input to slice of strings, i.e. []string
// Works for both csv (sep = ",") and row delimited (sep = "\n")
func ParseStringList(input string, sep string) []string {
	lines := strings.Split(input, sep)
	return lines
}

// Parse input to slice of strings, i.e. []string
// Works for both csv (sep = ",") and row delimited (sep = "\n")
func ParseStringStringList(input string, sep1 string, sep2 string) [][]string {
	var lists [][]string
	groups := strings.Split(input, sep1)
	for i := 0; i < len(groups); i++ {
		lists = append(lists, strings.Split(groups[i], sep2))
	}
	return lists
}

func ParseRuneListList(input string) [][]rune {
	var xys [][]rune
	for _, x := range strings.Split(input, "\n") {
		xys = append(xys, []rune(x))
	}
	return xys
}

func ParseRuneListListList(input string) [][][]rune {
	var xyzs [][][]rune
	var yzs [][]rune
	for _, x := range strings.Split(input, "\n") {
		if len(x) == 0 {
			xyzs = append(xyzs, yzs)
			yzs = nil
		} else {
			yzs = append(yzs, []rune(x))
		}
	}
	xyzs = append(xyzs, yzs)
	return xyzs
}
