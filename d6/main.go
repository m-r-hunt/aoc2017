package d6

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(6, main)
}

type state [16]int

var initialState = state{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6}

func main() (string, string) {
	visitedStates := map[state]int{}
	state := initialState
	visitedStates[state] = 0
	count := 0
	for {
		count++
		maxi := 0
		for i, mem := range state {
			if mem > state[maxi] {
				maxi = i
			}
		}
		redist := state[maxi]
		state[maxi] = 0

		for index := (maxi + 1) % 16; redist > 0; index = (index + 1) % 16 {
			state[index]++
			redist--
		}
		if _, ok := visitedStates[state]; ok {
			break
		}
		visitedStates[state] = count
	}

	return fmt.Sprint(count), fmt.Sprint(count - visitedStates[state])
}
