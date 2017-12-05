package main

import (
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"fmt"
)

func drawJumps(f *mygifs.Frame, jumps []int) {
	for i := 0; i < len(jumps)/2; i++ {
		avg := (jumps[i*2] + jumps[i*2+1]) / 20
		c := mygifs.ClampColour(mygifs.Colour(avg))
		f.SetPixel(i, 0, c)
		f.SetPixel(i, 1, c)
		f.SetPixel(i, 2, c)
		f.SetPixel(i, 3, c)
		f.SetPixel(i, 4, c)
	}
}

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	g := mygifs.NewGif(len(lines) / 2, 20)
	defer g.Write("solution.gif")
	g.SetFrameskip(1000)
	jumps := make([]int, len(lines))
	for i := range lines {
		n, _:= strconv.Atoi(lines[i])
		jumps[i] = n
	}
	pc := 0
	steps := 0
	for pc >= 0 && pc < len(jumps) {
		steps++
		npc := pc + jumps[pc]
		if jumps[pc] >= 3 {
			jumps[pc]--
		} else {
			jumps[pc]++
		}
		pc = npc
		f := g.AddBlankFrame()
		if f != nil {
			drawJumps(f, jumps)
			f.SetPixel(pc/2, 6, mygifs.Black)
		}
	}
	fmt.Println(steps)
}
