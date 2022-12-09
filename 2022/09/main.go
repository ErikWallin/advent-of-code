package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type Coordinate struct {
	x, y int
}

func isAdjacent(c1, c2 Coordinate) bool {
	return math.Abs(float64(c1.x-c2.x)) <= 1 && math.Abs(float64(c1.y-c2.y)) <= 1
}

func moveTail(head, tail *Coordinate) {
	if isAdjacent(*head, *tail) {
		return
	}
	if head.x == tail.x {
		if tail.y < head.y {
			tail.y++
		} else {
			tail.y--
		}
		return
	}
	if head.y == tail.y {
		if tail.x < head.x {
			tail.x++
		} else {
			tail.x--
		}
		return
	}
	if tail.y < head.y {
		tail.y++
	} else {
		tail.y--
	}
	if tail.x < head.x {
		tail.x++
	} else {
		tail.x--
	}
}

func run1(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	visited := mapset.NewSet[Coordinate]()
	hPos := Coordinate{0, 0}
	tPos := Coordinate{0, 0}
	visited.Add(tPos)
	for _, row := range list {
		items := strings.Split(row, " ")
		dir := items[0]
		amount := common.MustAtoi(items[1])
		for i := 0; i < amount; i++ {
			switch dir {
			case "L":
				hPos.x--
			case "R":
				hPos.x++
			case "U":
				hPos.y++
			case "D":
				hPos.y--
			}
			moveTail(&hPos, &tPos)
			visited.Add(tPos)
		}
	}
	//printMap(visited)
	return visited.Cardinality()
}

func run2(input string) interface{} {
	list := common.ParseStringList(input, "\n")
	visited := mapset.NewSet[Coordinate]()
	headPos := Coordinate{0, 0}
	tailPos := []Coordinate{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	visited.Add(tailPos[len(tailPos)-1])
	for _, row := range list {
		items := strings.Split(row, " ")
		dir := items[0]
		amount := common.MustAtoi(items[1])
		for a := 0; a < amount; a++ {
			for i := 0; i < len(tailPos); i++ {
				hPos := &headPos
				tPos := &tailPos[0]
				if i > 0 {
					hPos = &tailPos[i-1]
					tPos = &tailPos[i]
				}
				switch dir {
				case "L":
					if i == 0 {
						hPos.x--
					}
				case "R":
					if i == 0 {
						hPos.x++
					}
				case "U":
					if i == 0 {
						hPos.y++
					}
				case "D":
					if i == 0 {
						hPos.y--
					}
				}
				moveTail(hPos, tPos)
			}
			visited.Add(tailPos[len(tailPos)-1])
		}
	}

	//printMap(visited)

	return visited.Cardinality()
}

func printMap(visited mapset.Set[Coordinate]) {
	var minX, maxX, minY, maxY int
	for c := range visited.Iter() {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if visited.Contains(Coordinate{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	submit := true
	verbose := false
	common.Run(run1, run2, tests, 2022, 9, submit, verbose)
}
