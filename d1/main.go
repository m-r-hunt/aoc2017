package d1

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(1, part1, part2)
}

func drawBands(f *mygifs.Frame, s string) {
	for i, d := range s {
		for j := 6; j < 11; j++ {
			f.SetPixel(i, j, mygifs.Colour(1+d-'0'))
		}
	}
}

func drawLeftArrow(f *mygifs.Frame, i int, c mygifs.Colour) {

	f.SetPixel(i, 5, c)
	f.SetPixel(i-1, 4, c)
	f.SetPixel(i, 4, c)
	f.SetPixel(i, 3, c)
	f.SetPixel(i-1, 3, c)
	f.SetPixel(i-2, 3, c)
	f.SetPixel(i, 2, c)
	f.SetPixel(i, 1, c)
	f.SetPixel(i, 0, c)
}

func drawRightArrow(f *mygifs.Frame, i int, good bool) {
	colour := mygifs.Red
	if good {
		colour = mygifs.Green
	}

	f.SetPixel(i, 11, colour)
	f.SetPixel(i+1, 12, colour)
	f.SetPixel(i, 12, colour)
	f.SetPixel(i, 13, colour)
	f.SetPixel(i+1, 13, colour)
	f.SetPixel(i+2, 13, colour)
	f.SetPixel(i, 14, colour)
	f.SetPixel(i, 15, colour)
	f.SetPixel(i, 16, colour)
}

func main() {
	s := mygifs.JustLoadLines("input.txt")[0]
	mygifs.Delay = 1
	g := mygifs.NewGif(len(s), 50)
	//g.SetFrameskip(10)

	nextTotal := 0
	halfTotal := 0
	for i, d := range s {
		f := g.AddBlankFrame()
		drawBands(f, s)
		drawLeftArrow(f, i, mygifs.Colour(1+d-'0'))
		digitVal := int(d - '0')
		next := (i + 1) % len(s)
		if s[i] == s[next] {
			nextTotal += digitVal
			drawRightArrow(f, next, true)
		} else {
			drawRightArrow(f, next, false)
		}
		half := (i + len(s)/2) % len(s)
		if s[i] == s[half] {
			halfTotal += digitVal
			drawRightArrow(f, half, true)
		} else {
			drawRightArrow(f, half, false)
		}
		f.DrawText(next, 20, fmt.Sprintf("Part 1: %4v", nextTotal), mygifs.Black)
		f.DrawText(half, 20, fmt.Sprintf("Part 2: %4v", halfTotal), mygifs.Black)
	}
	fmt.Println(nextTotal)
	fmt.Println(halfTotal)
	g.Write("solution.gif")
}
