package main

import (
	"github.com/m-r-hunt/mygifs"
	"strings"
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func hexDist(x, y int) int {
	z := -x-y
	dist := (abs(x) + abs(y) + abs(z)) / 2
	return dist
}

func main() {
	dirs := strings.Split(mygifs.JustLoadLines("input.txt")[0], ",")
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
	fmt.Println(hexDist(x, y))
	fmt.Println(maxDist)
}
