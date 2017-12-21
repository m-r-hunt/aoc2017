package d3

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(3, main)
}

const (
	Up = iota
	Left
	Down
	Right
	MaxHeading
)

func addHeading(x, y, heading int) (int, int) {
	switch heading {
	case Up:
		return x, y - 1
	case Left:
		return x - 1, y
	case Down:
		return x, y + 1
	case Right:
		return x + 1, y
	default:
		panic("Bad heading")
	}
}

func main() (string, string) {
	// Part 1 via calculator/thought.
	ans1 := ""
	// Part
	origin := 6
	spaces := make([][]int, origin*2)
	for i := range spaces {
		spaces[i] = make([]int, origin*2)
	}
	x, y := origin, origin
	spaces[x][y] = 1
	heading := Down
	for spaces[x][y] < 277678 {
		leftTurn := (heading + 1) % MaxHeading
		if nx, ny := addHeading(x, y, leftTurn); spaces[nx][ny] == 0 {
			heading = leftTurn
			x, y = nx, ny
		} else {
			x, y = addHeading(x, y, heading)
		}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				spaces[x][y] += spaces[x+i][y+j]
			}
		}
	}
	return ans1, fmt.Sprint(spaces[x][y])
}
