package main

import (
	"fmt"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

func run1(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	x := []int{1}
	for _, row := range rows {
		if row == "noop" {
			x = append(x, x[len(x)-1])
			continue
		}
		x = append(x, x[len(x)-1], x[len(x)-1]+common.MustAtoi(strings.Split(row, " ")[1]))
	}
	is := []int{20, 60, 100, 140, 180, 220}
	res := 0
	for _, i := range is {
		res += i * x[i-1]
	}
	return res
}

func run2(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	xs := []int{1}
	for _, row := range rows {
		if row == "noop" {
			xs = append(xs, xs[len(xs)-1])
			continue
		}
		xs = append(xs, xs[len(xs)-1], xs[len(xs)-1]+common.MustAtoi(strings.Split(row, " ")[1]))
	}
	fmt.Printf("xs[]=%d\n", xs)
	for i := range xs[1:] {
		pos := i % 40
		if pos == 0 {
			fmt.Print("\n")
		}
		if pos >= xs[i]-1 && pos <= xs[i]+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
	return 0
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 10, submit, verbose)
}
