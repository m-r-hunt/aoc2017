package d25

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(25, main)
}

type state int

type instruction struct {
	write int
	move  int
	next  state
}

const (
	left  = -1
	right = 1
)

// Hand parsed...
var instructions = map[state]map[int]instruction{
	'A': {0: {1, right, 'B'}, 1: {0, left, 'B'}},
	'B': {0: {0, right, 'C'}, 1: {1, left, 'B'}},
	'C': {0: {1, right, 'D'}, 1: {0, left, 'A'}},
	'D': {0: {1, left, 'E'}, 1: {1, left, 'F'}},
	'E': {0: {1, left, 'A'}, 1: {0, left, 'D'}},
	'F': {0: {1, right, 'A'}, 1: {1, left, 'E'}},
}

func main() (string, string) {
	tape := map[int]int{}
	cursor := 0
	state := state('A')

	for n := 1; n <= 12629077; n++ {
		i := instructions[state][tape[cursor]]
		tape[cursor] = i.write
		cursor += i.move
		state = i.next
	}

	checksum := 0
	for _, v := range tape {
		checksum += v
	}

	return fmt.Sprint(checksum), "reboot printer"
}
