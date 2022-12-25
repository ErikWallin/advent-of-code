package main

import (
	"fmt"
	"math"

	"github.com/ErikWallin/advent-of-code/common"
)

func toDecimal(snafu string) int {
	numberSum := 0
	for p := len(snafu) - 1; p >= 0; p-- {
		number := 0
		if snafu[p] == '=' {
			number = -2
		} else if snafu[p] == '-' {
			number = -1
		} else {
			number = common.MustAtoi(string(snafu[p]))
		}
		numberSum += number * int(math.Pow(5, float64(len(snafu)-1-p)))
	}
	return numberSum
}

func run1(input string) interface{} {
	decimalSum := 0
	for _, row := range common.ParseStringList(input, "\n") {
		decimalSum += toDecimal(row)
	}
	remainder := decimalSum
	answer := ""
	for e := 19; e >= 0; e-- {
		positionValue := int(math.Pow(5, float64(e)))
		digit := (remainder+5*positionValue/2)/positionValue - 2
		if digit == -1 {
			answer += "-"
		} else if digit == -2 {
			answer += "="
		} else {
			answer += fmt.Sprintf("%d", digit)
		}
		remainder -= digit * positionValue
	}
	println("answer", answer)
	return 0
}

func run2(input string) interface{} {
	// No part 2 on 25th
	return 0
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 25, submit)
}
