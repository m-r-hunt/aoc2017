package d22

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
)

func init() {
	registry.RegisterDay(22, main)
}

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

type direction int

const (
	down direction = iota
	right
	up
	left
	max
)

func turn(dir direction, add int) direction {
	dir += direction(add)
	if dir < 0 {
		dir += max
	}
	dir = dir % max
	return dir
}

func move(x, y int, dir direction) (int, int) {
	switch dir {
	case down:
		y += 1
	case right:
		x += 1
	case up:
		y -= 1
	case left:
		x -= 1
	default:
		panic("bleh")
	}
	return x, y
}

func part1() int {
	origin := 1000
	lines := mygifs.JustLoadLines("d22/input.txt")
	grid := make([][]bool, origin*2)
	for i, _ := range grid {
		grid[i] = make([]bool, origin*2)
		if i-origin >= 0 && i-origin < len(lines) {
			for j, c := range lines[i-origin] {
				grid[i][j+origin] = (c == '#')
			}
		}
	}

	x, y := origin+len(lines[0])/2, origin+len(lines)/2
	dir := up

	infections := 0
	for n := 1; n <= 10000; n++ {
		if grid[y][x] {
			dir = turn(dir, -1)
		} else {
			dir = turn(dir, 1)
		}

		grid[y][x] = !grid[y][x]
		if grid[y][x] {
			infections++
		}

		x, y = move(x, y, dir)
	}

	return infections
}

func part2() int {
	origin := 1000
	lines := mygifs.JustLoadLines("d22/input.txt")
	grid := make([][]state, origin*2)
	for i, _ := range grid {
		grid[i] = make([]state, origin*2)
		if i-origin >= 0 && i-origin < len(lines) {
			for j, c := range lines[i-origin] {
				if c == '#' {
					grid[i][j+origin] = infected
				} else {
					grid[i][j+origin] = clean
				}
			}
		}
	}

	x, y := origin+len(lines[0])/2, origin+len(lines)/2
	dir := up

	infections := 0
	for n := 1; n <= 10000000; n++ {
		switch grid[y][x] {
		case infected:
			dir = turn(dir, -1)
		case clean:
			dir = turn(dir, 1)
		case flagged:
			dir = turn(dir, 2)
		case weakened:
			// Noop
		}

		grid[y][x] = (grid[y][x] + 1) % 4
		if grid[y][x] == infected {
			infections++
		}

		x, y = move(x, y, dir)
	}

	return infections
}

func main() (string, string) {
	return fmt.Sprint(part1()), fmt.Sprint(part2())
}
