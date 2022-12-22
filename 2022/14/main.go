package main

import (
	"fmt"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type Coordinate struct {
	x int
	y int
}

func run1(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	rocks := mapset.NewSet[Coordinate]()
	sands := mapset.NewSet[Coordinate]()
	xMin := 1000
	xMax := 0
	yMax := 0
	for _, r := range rows {
		rs := strings.Split(r, " -> ")
		c := Coordinate{-1, -1}
		for _, r := range rs {
			x := common.MustAtoi(strings.Split(r, ",")[0])
			if x < xMin {
				xMin = x
			}
			if x > xMax {
				xMax = x
			}
			y := common.MustAtoi(strings.Split(r, ",")[1])
			if y > yMax {
				yMax = y
			}
			if c.x == -1 {
				c = Coordinate{x, y}
				rocks.Add(c)
				continue
			}
			if x == c.x && y < c.y {
				for yi := y; yi <= c.y; yi++ {
					rocks.Add(Coordinate{x, yi})
				}
			} else if x == c.x && y > c.y {
				for yi := y; yi >= c.y; yi-- {
					rocks.Add(Coordinate{x, yi})
				}
			} else if y == c.y && x > c.x {
				for xi := x; xi >= c.x; xi-- {
					rocks.Add(Coordinate{xi, y})
				}
			} else if y == c.y && x < c.x {
				for xi := x; xi <= c.x; xi++ {
					rocks.Add(Coordinate{xi, y})
				}
			}
			c = Coordinate{x, y}
		}
	}

	//fmt.Printf("xMin=%d\n", xMin)
	//fmt.Printf("xMax=%d\n", xMax)
	fmt.Printf("yMax=%d\n", yMax)

	printMap(yMax, xMin, xMax, rocks, sands)

	for {
		sc := Coordinate{500, 0}
		for {
			scn := Coordinate{sc.x, sc.y + 1}
			if scn.y > yMax {
				return sands.Cardinality()
			}
			if !rocks.Contains(scn) && !sands.Contains(scn) {
				sc = scn
				continue
			}
			scnl := Coordinate{scn.x - 1, scn.y}
			if !rocks.Contains(scnl) && !sands.Contains(scnl) {
				sc = scnl
				continue
			}
			scnr := Coordinate{scn.x + 1, scn.y}
			if !rocks.Contains(scnr) && !sands.Contains(scnr) {
				sc = scnr
				continue
			}
			sands.Add(sc)
			break
		}
		//printMap(yMax, xMin, xMax, rocks, sands)
	}
}

func printMap(yMax int, xMin int, xMax int, rocks mapset.Set[Coordinate], sands mapset.Set[Coordinate]) {
	for y := 0; y <= yMax+2; y++ {
		for x := xMin; x <= xMax; x++ {
			c := Coordinate{x, y}
			if rocks.Contains(c) {
				fmt.Print("#")
			} else if sands.Contains(c) {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		println()
	}
}

func run2(input string) interface{} {
	rows := common.ParseStringList(input, "\n")
	rocks := mapset.NewSet[Coordinate]()
	sands := mapset.NewSet[Coordinate]()
	xMin := 1000
	xMax := 0
	yMax := 0
	for _, r := range rows {
		rs := strings.Split(r, " -> ")
		c := Coordinate{-1, -1}
		for _, r := range rs {
			x := common.MustAtoi(strings.Split(r, ",")[0])
			if x < xMin {
				xMin = x
			}
			if x > xMax {
				xMax = x
			}
			y := common.MustAtoi(strings.Split(r, ",")[1])
			if y > yMax {
				yMax = y
			}
			if c.x == -1 {
				c = Coordinate{x, y}
				rocks.Add(c)
				continue
			}
			if x == c.x && y < c.y {
				for yi := y; yi <= c.y; yi++ {
					rocks.Add(Coordinate{x, yi})
				}
			} else if x == c.x && y > c.y {
				for yi := y; yi >= c.y; yi-- {
					rocks.Add(Coordinate{x, yi})
				}
			} else if y == c.y && x > c.x {
				for xi := x; xi >= c.x; xi-- {
					rocks.Add(Coordinate{xi, y})
				}
			} else if y == c.y && x < c.x {
				for xi := x; xi <= c.x; xi++ {
					rocks.Add(Coordinate{xi, y})
				}
			}
			c = Coordinate{x, y}
		}
	}

	for x := 500 - yMax - 10; x < 500+yMax+10; x++ {
		rocks.Add(Coordinate{x, yMax + 2})
	}

	//fmt.Printf("xMin=%d\n", xMin)
	//fmt.Printf("xMax=%d\n", xMax)
	//fmt.Printf("yMax=%d\n", yMax)

	printMap(yMax, xMin, xMax, rocks, sands)

	for {
		sc := Coordinate{500, 0}
		for {
			scn := Coordinate{sc.x, sc.y + 1}
			if !rocks.Contains(scn) && !sands.Contains(scn) && scn.y < yMax+2 {
				sc = scn
				continue
			}
			scnl := Coordinate{scn.x - 1, scn.y}
			if !rocks.Contains(scnl) && !sands.Contains(scnl) {
				sc = scnl
				continue
			}
			scnr := Coordinate{scn.x + 1, scn.y}
			if !rocks.Contains(scnr) && !sands.Contains(scnr) {
				sc = scnr
				continue
			}
			sands.Add(sc)
			if sc.y == 0 {
				return sands.Cardinality()
			}
			break
		}
		//printMap(yMax, xMin, xMax, rocks, sands)
	}
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 14, submit)
}
