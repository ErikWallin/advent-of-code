package common

import (
	"strconv"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func ReverseSlice[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
