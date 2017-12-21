package d11

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strings"
)

func init() {
	registry.RegisterDay(11, main)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func hexDist(x, y int) int {
	z := -x - y
	dist := (abs(x) + abs(y) + abs(z)) / 2
	return dist
}

func main() (string, string) {
	dirs := strings.Split(mygifs.JustLoadLines("d11/input.txt")[0], ",")
	x := 0
	y := 0
	maxDist := 0
	for _, d := range dirs {
		switch d {
		case "n":
			y -= 1
		case "nw":
			x -= 1
		case "sw":
			x -= 1
			y += 1
		case "s":
			y += 1
		case "se":
			x += 1
		case "ne":
			x += 1
			y -= 1
		default:
			panic(d)
		}
		dist := hexDist(x, y)
		if dist > maxDist {
			maxDist = dist
		}
	}
	return fmt.Sprint(hexDist(x, y)), fmt.Sprint(maxDist)
}
