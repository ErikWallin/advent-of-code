package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) int {
	list := common.ParseStringList(input, "\n")
	sum := 0
	for _, row := range list {
		com1 := row[:len(row)/2]
		com2 := row[len(row)/2:]
		for _, c := range com1 {
			if strings.ContainsRune(com2, c) {
				if c >= 'a' {
					sum += 1 + int(c-'a')
					break
				} else {
					sum += 27 + int(c-'A')
					break
				}
			}
		}
	}
	return sum
}

func run2(input string) int {
	list := common.ParseStringList(input, "\n")
	sum := 0
	for i := 0; i < len(list)-1; i += 3 {
		for _, c := range list[i] {
			if strings.ContainsRune(list[i+1], c) && strings.ContainsRune(list[i+2], c) {
				if c >= 'a' {
					sum += 1 + int(c-'a')
					break
				} else {
					sum += 27 + int(c-'A')
					break
				}
			}
		}
	}
	return sum
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 3, submit, verbose)
}
