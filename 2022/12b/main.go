package main

import (
	"fmt"

	"github.com/ErikWallin/advent-of-code/common"
	"github.com/RyanCarrier/dijkstra"
)

func idString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func height(r rune) rune {
	if r == 'S' {
		return 'a'
	} else if r == 'E' {
		return 'z'
	} else {
		return r
	}
}

func run1(input string) interface{} {
	heightMap := common.ParseRuneListList(input)
	graph := dijkstra.NewGraph()
	var start, end int
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			vertexId := idString(x, y)
			graph.AddMappedVertex(vertexId)
			h := heightMap[y][x]
			if h == 'S' {
				id, _ := graph.GetMapping(vertexId)
				start = id
			} else if h == 'E' {
				id, _ := graph.GetMapping(vertexId)
				end = id
			}
		}
	}

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			p := idString(x, y)
			h := height(heightMap[y][x])
			if x+1 < len(heightMap[y]) {
				pn := idString(x+1, y)
				hn := height(heightMap[y][x+1])
				if h+1 >= hn {
					graph.AddMappedArc(p, pn, 1)
				}
				if hn+1 >= h {
					graph.AddMappedArc(pn, p, 1)
				}
			}
			if y+1 < len(heightMap) {
				pn := idString(x, y+1)
				hn := height(heightMap[y+1][x])
				if h+1 >= hn {
					graph.AddMappedArc(p, pn, 1)
				}
				if hn+1 >= h {
					graph.AddMappedArc(pn, p, 1)
				}
			}
		}
	}
	best, err := graph.Shortest(start, end)
	if err != nil {
		fmt.Printf("Could not find shortest distance: %v\n", err)
	}
	return int(best.Distance)
}

func run2(input string) interface{} {
	heightMap := common.ParseRuneListList(input)
	graph := dijkstra.NewGraph()
	starts := []int{}
	var end int
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			vertexId := idString(x, y)
			graph.AddMappedVertex(vertexId)
			h := heightMap[y][x]
			if h == 'S' || h == 'a' {
				id, _ := graph.GetMapping(vertexId)
				starts = append(starts, id)
			} else if h == 'E' {
				id, _ := graph.GetMapping(vertexId)
				end = id
			}
		}
	}
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			p := idString(x, y)
			h := height(heightMap[y][x])
			if x+1 < len(heightMap[y]) {
				pn := idString(x+1, y)
				hn := height(heightMap[y][x+1])
				if h+1 >= hn {
					graph.AddMappedArc(p, pn, 1)
				}
				if hn+1 >= h {
					graph.AddMappedArc(pn, p, 1)
				}
			}
			if y+1 < len(heightMap) {
				pn := idString(x, y+1)
				hn := height(heightMap[y+1][x])
				if h+1 >= hn {
					graph.AddMappedArc(p, pn, 1)
				}
				if hn+1 >= h {
					graph.AddMappedArc(pn, p, 1)
				}
			}
		}
	}

	record := 100000
	for _, start := range starts {
		best, err := graph.Shortest(start, end)
		if err == nil && int(best.Distance) < record {
			record = int(best.Distance)
		}
	}
	return record
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 12, submit)
}
