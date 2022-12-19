package main

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/ErikWallin/advent-of-code/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type Coordinate struct {
	x int
	y int
	z int
}

func run1(input string) interface{} {
	lavas := []Coordinate{}
	for _, row := range common.ParseStringList(input, "\n") {
		xyz := strings.Split(row, ",")
		lavas = append(lavas, Coordinate{common.MustAtoi(xyz[0]), common.MustAtoi(xyz[1]), common.MustAtoi(xyz[2])})
	}
	sum := 0
	for i, l := range lavas {
		for j, other := range lavas {
			if i == j {
				continue
			}
			same := 0
			diff1 := 0
			if l.x == other.x {
				same++
			}
			if math.Abs(float64(l.x-other.x)) == 1 {
				diff1++
			}
			if l.y == other.y {
				same++
			}
			if math.Abs(float64(l.y-other.y)) == 1 {
				diff1++
			}
			if l.z == other.z {
				same++
			}
			if math.Abs(float64(l.z-other.z)) == 1 {
				diff1++
			}
			if same == 2 && diff1 == 1 {
				sum++
			}
		}
	}
	return 6*len(lavas) - sum
}

var lavas mapset.Set[Coordinate]
var minX int
var maxX int
var minY int
var maxY int
var minZ int
var maxZ int

var okCache = mapset.NewSet[Coordinate]()
var failCache = mapset.NewSet[Coordinate]()

func run2(input string) interface{} {
	lavas = mapset.NewSet[Coordinate]()
	for _, row := range common.ParseStringList(input, "\n") {
		xyz := strings.Split(row, ",")
		lavas.Add(Coordinate{common.MustAtoi(xyz[0]), common.MustAtoi(xyz[1]), common.MustAtoi(xyz[2])})
	}
	minX = 100
	maxX = -100
	minY = 100
	maxY = -100
	minZ = 100
	maxZ = -100
	for _, l := range lavas.ToSlice() {
		if l.x < minX {
			minX = l.x
		}
		if l.x > maxX {
			maxX = l.x
		}
		if l.y < minY {
			minY = l.y
		}
		if l.y > maxY {
			maxY = l.y
		}
		if l.z < minZ {
			minZ = l.z
		}
		if l.z > maxZ {
			maxZ = l.z
		}
	}

	sum := 0
	for _, l := range lavas.ToSlice() {
		neighbours := []Coordinate{
			{l.x + 1, l.y, l.z},
			{l.x - 1, l.y, l.z},
			{l.x, l.y + 1, l.z},
			{l.x, l.y - 1, l.z},
			{l.x, l.y, l.z + 1},
			{l.x, l.y, l.z - 1},
		}
		for _, neighbour := range neighbours {
			if okCache.Contains(neighbour) || (!failCache.Contains(neighbour) && !lavas.Contains(neighbour) && wayOut(neighbour, mapset.NewSet[Coordinate](), 0)) {
				sum++
			}
		}
	}
	return sum
}

func wayOut(air Coordinate, history mapset.Set[Coordinate], iterations int) bool {
	history.Add(air)
	if air.x <= minX || air.x >= maxX || air.y <= minY || air.y >= maxY || air.z <= minZ || air.z >= maxZ {
		for _, h := range history.ToSlice() {
			okCache.Add(h)
		}
		return true
	}
	if iterations > 10000 {
		for _, h := range history.ToSlice() {
			failCache.Add(h)
		}
		return false
	}
	neighbours := []Coordinate{
		{air.x + 1, air.y, air.z},
		{air.x - 1, air.y, air.z},
		{air.x, air.y + 1, air.z},
		{air.x, air.y - 1, air.z},
		{air.x, air.y, air.z + 1},
		{air.x, air.y, air.z - 1},
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(neighbours), func(i, j int) { neighbours[i], neighbours[j] = neighbours[j], neighbours[i] })
	for _, neighbour := range neighbours {
		if okCache.Contains(neighbour) || (!failCache.Contains(neighbour) && !lavas.Contains(neighbour) && wayOut(neighbour, history.Clone(), iterations+1)) {
			for _, h := range history.ToSlice() {
				okCache.Add(h)
			}
			return true
		}
	}
	for _, h := range history.ToSlice() {
		failCache.Add(h)
	}
	return false
}

func main() {
	submit := false
	verbose := false
	common.Run(run1, run2, tests, 2022, 18, submit, verbose)
}
