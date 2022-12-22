package main

import (
	"math"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"

	mapset "github.com/deckarep/golang-set/v2"
)

type Coordinate struct {
	s rune
	x int
	y int
}

type Pair struct {
	sensor Coordinate
	beacon Coordinate
}

func run1(input string) interface{} {
	list := common.ParseStringList(input, "\n")

	goalRowY := 2000000
	if list[0] == "Sensor at x=2, y=18: closest beacon is at x=-2, y=15" {
		goalRowY = 10
	}
	goalRowSB := mapset.NewSet[int]()
	goalRowAir := mapset.NewSet[int]()

	pairs := []Pair{}
	for _, row := range list {
		parts := strings.Split(row, ": closest beacon is at ")
		sx := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[0], "Sensor at ", ""), ", ")[0], "=")[1])
		sy := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[0], "Sensor at ", ""), ", ")[1], "=")[1])
		s := Coordinate{'S', sx, sy}
		if sy == goalRowY {
			goalRowSB.Add(sx)
		}
		bx := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[1], "Sensor at ", ""), ", ")[0], "=")[1])
		by := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[1], "Sensor at ", ""), ", ")[1], "=")[1])
		b := Coordinate{'B', bx, by}
		if sy == goalRowY {
			goalRowSB.Add(bx)
		}
		pairs = append(pairs, Pair{s, b})
	}

	for _, p := range pairs {
		distance := int(math.Abs(float64(p.sensor.x-p.beacon.x)) + math.Abs(float64(p.sensor.y-p.beacon.y)))
		startX := p.sensor.x - (distance - int(math.Abs(float64(goalRowY-p.sensor.y))))
		endX := p.sensor.x + (distance - int(math.Abs(float64(goalRowY-p.sensor.y))))
		for x := startX + 1; x <= endX; x++ {
			if !goalRowSB.Contains(x) {
				goalRowAir.Add(x)
			}
		}
	}
	return goalRowAir.Cardinality()
}

func run2(input string) interface{} {
	list := common.ParseStringList(input, "\n")

	maxRowXY := 4000000
	if list[0] == "Sensor at x=2, y=18: closest beacon is at x=-2, y=15" {
		maxRowXY = 20
	}

	pairs := []Pair{}
	for _, row := range list {
		parts := strings.Split(row, ": closest beacon is at ")
		sx := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[0], "Sensor at ", ""), ", ")[0], "=")[1])
		sy := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[0], "Sensor at ", ""), ", ")[1], "=")[1])
		s := Coordinate{'S', sx, sy}
		bx := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[1], "Sensor at ", ""), ", ")[0], "=")[1])
		by := common.MustAtoi(strings.Split(strings.Split(strings.ReplaceAll(parts[1], "Sensor at ", ""), ", ")[1], "=")[1])
		b := Coordinate{'B', bx, by}
		pairs = append(pairs, Pair{s, b})
	}

	for y := 0; y <= maxRowXY; y++ {
		goalRowAir := []common.Range{}
		for _, p := range pairs {
			if p.sensor.y == y && p.sensor.x >= 0 && p.sensor.x <= maxRowXY {
				goalRowAir = common.AddToRange(goalRowAir, common.Range{From: p.sensor.x, To: p.sensor.x})
			}
			if p.beacon.y == y && p.beacon.x >= 0 && p.beacon.x <= maxRowXY {
				goalRowAir = common.AddToRange(goalRowAir, common.Range{From: p.beacon.x, To: p.beacon.x})
			}
			distance := int(math.Abs(float64(p.sensor.x-p.beacon.x)) + math.Abs(float64(p.sensor.y-p.beacon.y)))
			startX := p.sensor.x - distance + int(math.Abs(float64(y-p.sensor.y)))
			endX := p.sensor.x + distance - int(math.Abs(float64(y-p.sensor.y)))
			if endX > startX {
				goalRowAir = common.AddToRange(goalRowAir, common.Range{From: startX, To: endX})
			}
		}
		if len(goalRowAir) == 2 {
			return (goalRowAir[0].To+1)*4000000 + y
		}
	}

	return 0
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 15, submit)
}
