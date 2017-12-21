package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strings"
)

type rule struct {
	input  string
	output string
}

func makeGrid(size int) [][]bool {
	g := make([][]bool, size)
	for i := range g {
		g[i] = make([]bool, size)
	}
	return g
}

func applyRule2(i, j int, grid *[][]bool, r rule) {
	i = i / 2 * 3
	j = j / 2 * 3
	fs := strings.Split(r.output, "/")
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			(*grid)[i+m][j+n] = fs[m][n] == '#'
		}
	}
}

func enhance2(i, j int, oldGrid, newGrid *[][]bool, rules []rule) {
	for _, r := range rules {
		fs := strings.Split(r.input, "/")

		found := true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[m][n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[m][n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[1-m][n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[1-m][n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[m][1-n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[m][1-n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[1-m][1-n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[1-m][1-n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[n][m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[n][m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[n][1-m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[n][1-m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[1-n][m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[1-n][m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 2; m++ {
			for n := 0; n < 2; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[1-n][1-m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[1-n][1-m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule2(i, j, newGrid, r)
			return
		}
	}
	panic("No rule matched")
}

func applyRule3(i, j int, grid *[][]bool, r rule) {
	i = i / 3 * 4
	j = j / 3 * 4
	fs := strings.Split(r.output, "/")
	for m := 0; m < 4; m++ {
		for n := 0; n < 4; n++ {
			(*grid)[i+m][j+n] = fs[m][n] == '#'
		}
	}
}

func enhance3(i, j int, oldGrid, newGrid *[][]bool, rules []rule) {
	for _, r := range rules {
		fs := strings.Split(r.input, "/")

		found := true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[m][n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[m][n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[2-m][n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[2-m][n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[m][2-n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[m][2-n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[2-m][2-n] == '.') || (!(*oldGrid)[i+m][j+n] && fs[2-m][2-n] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[n][m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[n][m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[n][2-m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[n][2-m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[2-n][m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[2-n][m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}

		found = true
		for m := 0; m < 3; m++ {
			for n := 0; n < 3; n++ {
				if ((*oldGrid)[i+m][j+n] && fs[2-n][2-m] == '.') || (!(*oldGrid)[i+m][j+n] && fs[2-n][2-m] == '#') {
					found = false
					break
				}
			}
		}
		if found {
			applyRule3(i, j, newGrid, r)
			return
		}
	}
	panic("No rule matched")
}

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	rules2 := []rule{}
	rules3 := []rule{}
	for _, l := range lines {
		f := strings.Fields(l)
		if len(f[0]) == 5 {
			rules2 = append(rules2, rule{f[0], f[2]})
		} else {
			rules3 = append(rules3, rule{f[0], f[2]})
		}
	}
	grid := [][]bool{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	for n := 0; n < 18; n++ {
		if len(grid)%2 == 0 {
			newGrid := makeGrid(len(grid) + len(grid)/2)
			for i := 0; i < len(grid); i += 2 {
				for j := 0; j < len(grid); j += 2 {
					enhance2(i, j, &grid, &newGrid, rules2)
				}
			}
			grid = newGrid
		} else if len(grid)%3 == 0 {
			newGrid := makeGrid(len(grid) + len(grid)/3)
			for i := 0; i < len(grid); i += 3 {
				for j := 0; j < len(grid); j += 3 {
					enhance3(i, j, &grid, &newGrid, rules3)
				}
			}
			grid = newGrid
		} else {
			panic("Bad grid size")
		}
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
