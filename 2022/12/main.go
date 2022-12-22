package main

import (
	"github.com/ErikWallin/advent-of-code/common"
)

func id(x, y int) int {
	return 1000000 + x*10000 + y
}

func run1(input string) interface{} {
	heightMap := common.ParseRuneListList(input)
	changed := true
	distances := map[int]int{}
	var start int
	var end int
	for changed {
		changed = false
		for y := 0; y < len(heightMap); y++ {
			for x := 0; x < len(heightMap[0]); x++ {
				p := id(x, y)
				h := heightMap[y][x]
				if heightMap[y][x] == 'S' {
					h = 'a'
					start = p
					distances[start] = 0
				}
				if heightMap[y][x] == 'E' {
					h = 'z'
					end = p
				}
				d, ok := distances[p]
				if !ok {
					continue
				}
				// right
				if x+1 < len(heightMap[0]) {
					pn := id(x+1, y)
					hn := heightMap[y][x+1]
					if start == pn {
						hn = 'a'
					} else if end == pn {
						hn = 'z'
					}
					if hn-1 <= h {
						dn, okn := distances[pn]
						if !okn {
							distances[pn] = d + 1
							changed = true
						} else if dn > d+1 {
							distances[pn] = d + 1
							changed = true
						}
					}
				}
				// left
				if x-1 >= 0 {
					pn := id(x-1, y)
					hn := heightMap[y][x-1]
					if start == pn {
						hn = 'a'
					} else if end == pn {
						hn = 'z'
					}
					if hn-1 <= h {
						dn, okn := distances[pn]
						if !okn {
							distances[pn] = d + 1
							changed = true
						} else if dn > d+1 {
							distances[pn] = d + 1
							changed = true
						}
					}
				}
				// up
				if y-1 >= 0 {
					pn := id(x, y-1)
					hn := heightMap[y-1][x]
					if start == pn {
						hn = 'a'
					} else if end == pn {
						hn = 'z'
					}
					if hn-1 <= h {
						dn, okn := distances[pn]
						if !okn {
							distances[pn] = d + 1
							changed = true
						} else if dn > d+1 {
							distances[pn] = d + 1
							changed = true
						}
					}
				}
				// down
				if y+1 < len(heightMap) {
					pn := id(x, y+1)
					hn := heightMap[y+1][x]
					if start == pn {
						hn = 'a'
					} else if end == pn {
						hn = 'z'
					}
					if hn-1 <= h {
						dn, okn := distances[pn]
						if !okn {
							distances[pn] = d + 1
							changed = true
						} else if dn > d+1 {
							distances[pn] = d + 1
							changed = true
						}
					}
				}
			}
		}
	}
	return distances[end]
}

func run2(input string) interface{} {
	heightMap := common.ParseRuneListList(input)

	var end int
	best := 100000

	var starts []int
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[0]); x++ {
			if heightMap[y][x] == 'a' || heightMap[y][x] == 'S' {
				starts = append(starts, id(x, y))
			}
		}
	}

	for _, start := range starts {
		changed := true
		distances := map[int]int{}
		distances[start] = 0
		for changed {
			changed = false
			for y := 0; y < len(heightMap); y++ {
				for x := 0; x < len(heightMap[0]); x++ {
					p := id(x, y)
					h := heightMap[y][x]
					if heightMap[y][x] == 'S' {
						h = 'a'
					}
					if heightMap[y][x] == 'E' {
						h = 'z'
						end = p
					}
					d, ok := distances[p]
					if !ok {
						continue
					}
					// right
					if x+1 < len(heightMap[0]) {
						pn := id(x+1, y)
						hn := heightMap[y][x+1]
						if start == pn {
							hn = 'a'
						} else if end == pn {
							hn = 'z'
						}
						if hn-1 <= h {
							dn, okn := distances[pn]
							if !okn {
								distances[pn] = d + 1
								changed = true
							} else if dn > d+1 {
								distances[pn] = d + 1
								changed = true
							}
						}
					}
					// left
					if x-1 >= 0 {
						pn := id(x-1, y)
						hn := heightMap[y][x-1]
						if start == pn {
							hn = 'a'
						} else if end == pn {
							hn = 'z'
						}
						if hn-1 <= h {
							dn, okn := distances[pn]
							if !okn {
								distances[pn] = d + 1
								changed = true
							} else if dn > d+1 {
								distances[pn] = d + 1
								changed = true
							}
						}
					}
					// up
					if y-1 >= 0 {
						pn := id(x, y-1)
						hn := heightMap[y-1][x]
						if start == pn {
							hn = 'a'
						} else if end == pn {
							hn = 'z'
						}
						if hn-1 <= h {
							dn, okn := distances[pn]
							if !okn {
								distances[pn] = d + 1
								changed = true
							} else if dn > d+1 {
								distances[pn] = d + 1
								changed = true
							}
						}
					}
					// down
					if y+1 < len(heightMap) {
						pn := id(x, y+1)
						hn := heightMap[y+1][x]
						if start == pn {
							hn = 'a'
						} else if end == pn {
							hn = 'z'
						}
						if hn-1 <= h {
							dn, okn := distances[pn]
							if !okn {
								distances[pn] = d + 1
								changed = true
							} else if dn > d+1 {
								distances[pn] = d + 1
								changed = true
							}
						}
					}
				}
			}
		}
		d, ok := distances[end]
		if !ok {
			continue
		} else if best > d {
			best = d
		} else {
		}
	}
	return best
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 12, submit)
}
