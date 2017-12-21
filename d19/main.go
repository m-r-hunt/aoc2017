package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ttype int

const (
	straight ttype = iota
	corner
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

func main() {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(0)
	}
	lines := strings.Split(string(d), "\n")
	for i, v := range lines {
		lines[i] = strings.Trim(v, "\n\r\t")
	}

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[0:len(lines)-1]
	}
	grid := make([][]ttype, len(lines))
	startx := 0
	for i, l := range lines {
		grid[i] = make([]ttype, len(l))
		for j, c := range l {
			switch c {
			case '|':
				if i == 0 {
					startx = j
				}
				fallthrough
			case '-':
				grid[i][j] = straight
			case '+':
				grid[i][j] = corner
			default:
				grid[i][j] = ttype(c)
			}
		}
	}
	fmt.Println(startx)
	fmt.Println(grid)
	out := ""
	x, y := startx, 0
	dir := down
	count := 0
	loop:
	for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
		switch grid[y][x] {
		default:
			out += fmt.Sprintf("%c", grid[y][x])
			fallthrough
		case straight:
			x, y = move(x, y, dir)
		case corner:
			x1, y1 := move(x, y, turn(dir, -1))
			x2, y2 := move(x, y, turn(dir, 1))
			if x1 >= 0 && x1 < len(grid[0]) && y1 >= 0 && y1 < len(grid)&& grid[y1][x1] != ' ' {
				x, y = x1, y1
				dir = turn(dir, -1)
			} else if x2 >= 0 && x2 < len(grid[0]) && y2 >= 0 && y2 < len(grid) && grid[y2][x2] != ' ' {
				x, y = x2, y2
				dir = turn(dir, 1)
			} else {
				panic("fblewh")
		}
		case ' ':
			break loop
		}
		count++
	}
	fmt.Println(out)
	fmt.Println(count)
}
