package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strings"
)

type rule struct {
	input  []string
	output []string
}

func makeGrid(size int) [][]bool {
	g := make([][]bool, size)
	for i := range g {
		g[i] = make([]bool, size)
	}
	return g
}

func applyRule(size int, i, j int, grid *[][]bool, r rule) {
	i = i / size * (size + 1)
	j = j / size * (size + 1)
	for m := 0; m < size+1; m++ {
		for n := 0; n < size+1; n++ {
			(*grid)[i+m][j+n] = (r.output[m][n] == '#')
		}
	}
}

func enhance(size int, i, j int, oldGrid [][]bool, newGrid *[][]bool, rules []rule) {

	transforms := []func(m, n int) (int, int){
		func(m, n int) (int, int) { return m, n },
		func(m, n int) (int, int) { return size - 1 - m, n },
		func(m, n int) (int, int) { return m, size - 1 - n },
		func(m, n int) (int, int) { return size - 1 - m, size - 1 - n },
		func(m, n int) (int, int) { return n, m },
		func(m, n int) (int, int) { return size - 1 - n, m },
		func(m, n int) (int, int) { return n, size - 1 - m },
		func(m, n int) (int, int) { return size - 1 - n, size - 1 - m },
	}

	tryTransformedRule := func(t func(m, n int) (int, int), r rule) bool {
		for m := 0; m < size; m++ {
			for n := 0; n < size; n++ {
				mt, nt := t(m, n)
				if oldGrid[i+m][j+n] != (r.input[mt][nt] == '#') {
					return false
				}
			}
		}
		return true
	}

	for _, r := range rules {
		for _, t := range transforms {
			if tryTransformedRule(t, r) {
				applyRule(size, i, j, newGrid, r)
				return
			}
		}
	}

	panic("No rule matched")
}

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	rules := map[int][]rule{2: []rule{}, 3: []rule{}}
	for _, l := range lines {
		f := strings.Fields(l)
		r := rule{strings.Split(f[0], "/"), strings.Split(f[2], "/")}
		if len(r.input) == 2 {
			rules[2] = append(rules[2], r)
		} else {
			rules[3] = append(rules[3], r)
		}
	}

	grid := [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	for n := 0; n < 18; n++ {
		size := 3
		if len(grid)%2 == 0 {
			size = 2
		}

		newGrid := makeGrid(len(grid) + len(grid)/size)
		for i := 0; i < len(grid); i += size {
			for j := 0; j < len(grid); j += size {
				enhance(size, i, j, grid, &newGrid, rules[size])
			}
		}
		grid = newGrid
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
