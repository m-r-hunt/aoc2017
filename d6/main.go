package main

import (
	"fmt"
)

type state [16]int

var initialState = state{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6}

func main() {
	visitedStates := map[state]int{}
	state := initialState
	visitedStates[state] = 0
	count := 0
	fmt.Println(state)
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
		fmt.Println(state)
		if _, ok := visitedStates[state]; ok {
			break
		}
		visitedStates[state] = count
	}
	fmt.Println(count)
	fmt.Println(count - visitedStates[state])
}
