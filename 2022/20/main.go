package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

type IndexedValue struct {
	index int
	value int
}

func run1(input string) interface{} {
	ints := common.ParseIntList(input, "\n")
	mixed := make([]IndexedValue, len(ints))
	for i, v := range ints {
		mixed[i] = IndexedValue{i, v}
	}
	for originIndex := 0; originIndex < len(mixed); originIndex++ {
		iCurrent := indexOfIndex(mixed, originIndex)
		current := mixed[iCurrent]
		move := (current.value + 2*(len(mixed)-1)) % (len(mixed) - 1)
		if iCurrent+move >= len(mixed) {
			move -= (len(mixed) - 1)
		}
		if move == 0 {
			//fmt.Printf("%d no change\n", originIndex)
		} else if move < 0 {
			copy(mixed[iCurrent+move+1:], mixed[iCurrent+move:iCurrent])
			mixed[iCurrent+move] = current
		} else {
			copy(mixed[iCurrent:], mixed[iCurrent+1:iCurrent+move+1])
			mixed[iCurrent+move] = current
		}
	}
	i0 := indexOfValue(mixed, 0)
	v1000 := mixed[(i0+1000)%len(mixed)].value
	v2000 := mixed[(i0+2000)%len(mixed)].value
	v3000 := mixed[(i0+3000)%len(mixed)].value
	return v1000 + v2000 + v3000
}

func indexOfIndex(slice []IndexedValue, index int) int {
	for i, iv := range slice {
		if iv.index == index {
			return i
		}
	}
	return -1
}

func indexOfValue(slice []IndexedValue, value int) int {
	for i, iv := range slice {
		if iv.value == value {
			return i
		}
	}
	return -1
}

func run2(input string) interface{} {
	decryptionKey := 811589153
	ints := common.ParseIntList(input, "\n")
	mixed := make([]IndexedValue, len(ints))
	for i, v := range ints {
		mixed[i] = IndexedValue{i, v * decryptionKey}
	}
	for r := 0; r < 10; r++ {
		for originIndex := 0; originIndex < len(mixed); originIndex++ {
			iCurrent := indexOfIndex(mixed, originIndex)
			current := mixed[iCurrent]
			move := (current.value + 2*decryptionKey*(len(mixed)-1)) % (len(mixed) - 1)
			if iCurrent+move >= len(mixed) {
				move -= (len(mixed) - 1)
			}
			if move == 0 {
				//fmt.Printf("%d no change\n", originIndex)
			} else if move < 0 {
				copy(mixed[iCurrent+move+1:], mixed[iCurrent+move:iCurrent])
				mixed[iCurrent+move] = current
			} else {
				copy(mixed[iCurrent:], mixed[iCurrent+1:iCurrent+move+1])
				mixed[iCurrent+move] = current
			}
		}
	}
	i0 := indexOfValue(mixed, 0)
	v1000 := mixed[(i0+1000)%len(mixed)].value
	v2000 := mixed[(i0+2000)%len(mixed)].value
	v3000 := mixed[(i0+3000)%len(mixed)].value
	return v1000 + v2000 + v3000
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 20, submit, verbose)
}
