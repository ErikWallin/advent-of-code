package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	list := common.ParseStringList(input, "\n")[1:]
	var stack []string
	dirs := map[string]int{}
	shouldAdd := true
	for _, row := range list {
		if row == "$ cd .." {
			stack = stack[:len(stack)-1]
			continue
		}
		args := strings.Split(row, " ")
		if args[0] == "dir" || args[1] == "ls" {
			continue
		}
		if args[1] == "cd" {
			stack = append(stack, args[2])
			dirs[path(stack)] = 0
			continue
		}
		size := common.MustAtoi(args[0])
		if shouldAdd {
			for d, _ := range stack {
				p := path(stack[:d+1])
				if _, ok := dirs[p]; !ok {
					dirs[p] = 0
				}
				dirs[p] = dirs[p] + size
			}
		}
	}
	sum := 0
	for _, v := range dirs {
		if v < 100000 {
			sum += v
		}
	}
	return sum
}

func path(stack []string) string {
	path := ""
	for _, s := range stack {
		path += "/" + s
	}
	return path
}

func run2(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	var stack []string
	dirs := map[string]int{}
	shouldAdd := true
	for _, row := range list {
		if row == "$ cd .." {
			stack = stack[:len(stack)-1]
			continue
		}
		args := strings.Split(row, " ")
		if args[0] == "dir" || args[1] == "ls" {
			continue
		}
		if args[1] == "cd" {
			stack = append(stack, args[2])
			dirs[path(stack)] = 0
			continue
		}
		size := common.MustAtoi(args[0])
		if shouldAdd {
			for d, _ := range stack {
				p := path(stack[:d+1])
				if _, ok := dirs[p]; !ok {
					dirs[p] = 0
				}
				dirs[p] = dirs[p] + size
			}
		}
	}
	unused := 70000000 - dirs["//"]
	required := 30000000 - unused
	smallest := 70000000
	for _, v := range dirs {
		if v > required && v < smallest {
			smallest = v
		}
	}
	return smallest
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 7, submit)
}
