package main

import (
	"strings"

	"github.com/ErikWallin/advent-of-code/common"
)

type Position struct {
	x      int
	y      int
	facing rune
}

func finalPassword(position Position) int {
	sum := (position.y+1)*1000 + (position.x+1)*4
	if position.facing == '>' {
		sum += 0
	} else if position.facing == '^' {
		sum += 3
	} else if position.facing == '<' {
		sum += 2
	} else if position.facing == 'v' {
		sum += 1
	}
	return sum
}

func run1(input string) interface{} {
	groups := common.ParseStringList(input, "\n\n")
	path := groups[1]
	field := common.ParseRuneListList(groups[0])
	position := Position{strings.Index(string(field[0]), "."), 0, '>'}
	currentNumber := ""
	for _, c := range path {
		if c >= '0' && c <= '9' {
			currentNumber += string(c)
		} else {
			if currentNumber != "" {
				position = forward1(field, position, position, common.MustAtoi(currentNumber))
			}
			currentNumber = ""
			position = turn(position, c)
		}
	}
	if currentNumber != "" {
		position = forward1(field, position, position, common.MustAtoi(currentNumber))
	}
	return finalPassword(position)
}

func forward1(field [][]rune, lastOpenPosition Position, position Position, amount int) Position {
	if amount == 0 {
		return Position{lastOpenPosition.x, lastOpenPosition.y, lastOpenPosition.facing}
	}
	nextX := position.x
	nextY := position.y
	if position.facing == '<' {
		nextX--
		if nextX < 0 {
			nextX = len(field[nextY]) - 1
		}
	} else if position.facing == '>' {
		nextX++
		if nextX >= len(field[nextY]) {
			nextX = 0
		}
	} else if position.facing == 'v' {
		nextY++
		if nextY >= len(field) {
			nextY = 0
		}
	} else if position.facing == '^' {
		nextY--
		if nextY < 0 {
			nextY = len(field) - 1
		}
	}
	if len(field[nextY]) < nextX+1 || field[nextY][nextX] == ' ' {
		return forward1(field, lastOpenPosition, Position{nextX, nextY, position.facing}, amount)
	} else if field[nextY][nextX] == '#' {
		return Position{lastOpenPosition.x, lastOpenPosition.y, lastOpenPosition.facing}
	} else {
		newLastOpenPosition := Position{nextX, nextY, position.facing}
		return forward1(field, newLastOpenPosition, newLastOpenPosition, amount-1)
	}
}

func turn(position Position, where rune) Position {
	var facing rune
	if where == 'L' {
		if position.facing == '>' {
			facing = '^'
		} else if position.facing == '^' {
			facing = '<'
		} else if position.facing == '<' {
			facing = 'v'
		} else if position.facing == 'v' {
			facing = '>'
		}
	} else if where == 'R' {
		if position.facing == '>' {
			facing = 'v'
		} else if position.facing == '^' {
			facing = '>'
		} else if position.facing == '<' {
			facing = '^'
		} else if position.facing == 'v' {
			facing = '<'
		}
	} else {
		panic("unknown turn")
	}
	return Position{position.x, position.y, facing}
}

func run2(input string) interface{} {
	groups := common.ParseStringList(input, "\n\n")
	path := groups[1]
	if path == "10R5L5R10L4R5L5" {
		initConnectionsTest()
	} else {
		initConnectionsProd()
	}
	field := common.ParseRuneListList(groups[0])
	position := Position{strings.Index(string(field[0]), "."), 0, '>'}
	currentNumber := ""
	for _, c := range path {
		if c >= '0' && c <= '9' {
			currentNumber += string(c)
		} else {
			if currentNumber != "" {
				position = forward2(field, position, common.MustAtoi(currentNumber))
			}
			currentNumber = ""
			position = turn(position, c)
		}
	}
	if currentNumber != "" {
		position = forward2(field, position, common.MustAtoi(currentNumber))
	}
	return finalPassword(position)
}

func forward2(field [][]rune, position Position, amount int) Position {
	if amount == 0 {
		return Position{position.x, position.y, position.facing}
	}
	nextPosition := nextPos(field, position)
	if field[nextPosition.y][nextPosition.x] == '#' {
		return Position{position.x, position.y, position.facing}
	} else {
		return forward2(field, nextPosition, amount-1)
	}
}

func nextPos(field [][]rune, position Position) Position {
	nextPosition, ok := connections[position]
	if ok {
		return nextPosition
	}
	nextX := position.x
	nextY := position.y
	if position.facing == '<' {
		nextX--
	} else if position.facing == '>' {
		nextX++
	} else if position.facing == 'v' {
		nextY++
	} else if position.facing == '^' {
		nextY--
	}
	return Position{nextX, nextY, position.facing}
}

var connections map[Position]Position

func initConnectionsTest() {
	connections = map[Position]Position{}
	for i := 0; i < 4; i++ {
		connections[Position{4 + i, 4, '^'}] = Position{8, i, '>'}
		connections[Position{8, i, '<'}] = Position{4 + i, 4, 'v'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{i, 4, '^'}] = Position{11 - i, 0, 'v'}
		connections[Position{11 - i, 0, '^'}] = Position{i, 4, 'v'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{0, 4 + i, '<'}] = Position{15 - i, 11, '^'}
		connections[Position{15 - i, 11, 'v'}] = Position{0, 4 + i, '>'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{i, 7, 'v'}] = Position{11 - i, 11, '^'}
		connections[Position{11 - i, 11, 'v'}] = Position{i, 7, '^'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{7, 4 + i, 'v'}] = Position{8, 11 - i, '>'}
		connections[Position{8, 11 - i, '<'}] = Position{7, 4 + i, '^'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{7, 4 + i, 'v'}] = Position{8, 11 - i, '>'}
		connections[Position{8, 11 - i, '<'}] = Position{7, 4 + i, '^'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{11, 4 + i, '>'}] = Position{15 - i, 8, 'v'}
		connections[Position{15 - i, 8, '^'}] = Position{11, 4 + i, '<'}
	}
	for i := 0; i < 4; i++ {
		connections[Position{i, 11, '>'}] = Position{15, 11 - i, '<'}
		connections[Position{15, 11 - i, '>'}] = Position{i, 11, '<'}
	}
}

func initConnectionsProd() {
	connections = map[Position]Position{}
	for i := 0; i < 50; i++ {
		connections[Position{i, 100, '^'}] = Position{50, 50 + i, '>'}
		connections[Position{50, 50 + i, '<'}] = Position{i, 100, 'v'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{0, 100 + i, '<'}] = Position{50, 49 - i, '>'}
		connections[Position{50, 49 - i, '<'}] = Position{0, 100 + i, '>'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{0, 150 + i, '<'}] = Position{50 + i, 0, 'v'}
		connections[Position{50 + i, 0, '^'}] = Position{0, 150 + i, '>'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{i, 199, 'v'}] = Position{100 + i, 0, 'v'}
		connections[Position{100 + i, 0, '^'}] = Position{i, 199, '^'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{49, 150 + i, '>'}] = Position{50 + i, 149, '^'}
		connections[Position{50 + i, 149, 'v'}] = Position{49, 150 + i, '<'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{99, 100 + i, '>'}] = Position{149, 49 - i, '<'}
		connections[Position{149, 49 - i, '>'}] = Position{99, 100 + i, '<'}
	}
	for i := 0; i < 50; i++ {
		connections[Position{99, 50 + i, '>'}] = Position{100 + i, 49, '^'}
		connections[Position{100 + i, 49, 'v'}] = Position{99, 50 + i, '<'}
	}
}

func main() {
	submit := false
	common.Run(run1, run2, tests, 2022, 22, submit)
}
