package common

import (
	"sort"
)

type Range struct {
	From int
	To   int
}

func AddToRange(ranges []Range, r Range) []Range {
	if len(ranges) == 0 {
		return []Range{r}
	}
	ranges = append(ranges, r)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})
	res := []Range{}
	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i].To >= ranges[i+1].From-1 {
			res = append(res, Range{ranges[i].From, MaxOfInts(ranges[i].To, ranges[i+1].To)})
			for j := i + 2; j < len(ranges); j++ {
				res = AddToRange(res, ranges[j])
			}
			return res
		} else {
			res = append(res, ranges[i])
			if i == len(ranges)-2 {
				res = append(res, ranges[i+1])
			}
		}
	}
	return res
}

func MaxOfInts(vars ...int) int {
	if len(vars) == 0 {
		var zero int
		return zero
	}
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func MinOfInts(vars ...int) int {
	if len(vars) == 0 {
		var zero int
		return zero
	}
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

/*
func MinOf[T constraints.Ordered](vars ...T) T {
	if len(vars) == 0 {
		var zero T
		return zero
	}
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}*/
