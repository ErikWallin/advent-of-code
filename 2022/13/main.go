package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

type Component interface {
	compare(other Component) int
	isLeaf() bool
}

type Packet struct {
	components []Component
}

func (p *Packet) Add(c Component) {
	p.components = append(p.components, c)
}

func (p Packet) compare(other Component) int {
	if other.isLeaf() {
		np := Packet{[]Component{}}
		np.Add(other)
		return p.compare(np)
	}
	otherPacket := other.(Packet)
	otherPacketLen := len(otherPacket.components)
	for i, c := range p.components {
		if i >= otherPacketLen {
			return 1
		}
		o := otherPacket.components[i]
		cmp := c.compare(o)
		if cmp == 0 {
			continue
		} else {
			return cmp
		}
	}
	return len(p.components) - otherPacketLen
}

func (p Packet) isLeaf() bool {
	return false
}

type PacketLeaf struct {
	val int
}

func (p PacketLeaf) compare(other Component) int {
	if other.isLeaf() {
		return p.val - other.(PacketLeaf).val
	} else {
		np := Packet{[]Component{}}
		np.Add(p)
		return np.compare(other)
	}
}

func (p PacketLeaf) isLeaf() bool {
	return true
}

func parse(s string) Component {
	i, err := strconv.Atoi(s)
	if err == nil {
		return PacketLeaf{i}
	}
	s = s[1 : len(s)-1]
	p := Packet{[]Component{}}
	var groups [][]rune
	var current []rune
	stack := 0
	for _, r := range s {
		if r == ',' && stack == 0 {
			groups = append(groups, current)
			current = []rune{}
			continue
		}
		if r == '[' {
			stack++
		}
		if r == ']' {
			stack--
		}
		current = append(current, r)
	}
	if current != nil {
		groups = append(groups, current)
	}
	for _, g := range groups {
		p.components = append(p.components, parse(string(g)))
	}
	return p
}

func run1(input string) interface{} {
	groups := common.ParseStringStringList(input, "\n\n", "\n")
	sum := 0
	for i, g := range groups {
		p1 := parse(g[0])
		p2 := parse(g[1])
		if p1.compare(p2) <= 0 {
			sum += i + 1
		}
	}
	return sum
}

func run2(input string) interface{} {
	packetStrings := common.ParseStringList(strings.ReplaceAll(input, "\n\n", "\n"), "\n")
	components := []Component{}
	for _, s := range packetStrings {
		components = append(components, parse(s))
	}
	c2 := parse("[[2]]")
	c6 := parse("[[6]]")
	components = append(components, c2, c6)

	sort.Slice(components, func(i, j int) bool {
		return components[i].compare(components[j]) < 0
	})

	decoderKeyIndecies := []int{}
	for i, c := range components {
		s := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", c), "{", ""), "}", "")
		if s == "[[2]]" || s == "[[6]]" {
			decoderKeyIndecies = append(decoderKeyIndecies, i+1)
		}
	}
	return decoderKeyIndecies[0] * decoderKeyIndecies[1]
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 13, submit)
}
