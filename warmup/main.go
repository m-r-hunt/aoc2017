package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	mygifs.Delay = 1
	g := mygifs.NewGif(200, 200)
	g.SetCamera(100, 100)
	s := mygifs.JustLoadLines("input.txt")
	i := strings.Split(s[0], ", ")
	x, y := 0, 0
	aMarkers, bMarkers := []struct{ x, y int }{}, []struct{ x, y int }{}
	var f *mygifs.Frame
	for i, s := range i {
		if i%50 == 0 || i == len(s)-1 {
			f = g.AddCopyFrame()
		}
		switch s {
		case "Up":
			y -= 1
			f.SetPixel(x, y, mygifs.Black)
		case "Down":
			y += 1
			f.SetPixel(x, y, mygifs.Black)
		case "Left":
			x -= 1
			f.SetPixel(x, y, mygifs.Black)
		case "Right":
			x += 1
			f.SetPixel(x, y, mygifs.Black)
		case "A":
			aMarkers = append(aMarkers, struct{ x, y int }{x, y})
		case "B":
			bMarkers = append(bMarkers, struct{ x, y int }{x, y})
		}
		for _, m := range aMarkers {
			f.SetPixel(m.x, m.y, mygifs.Red)
		}
		for _, m := range bMarkers {
			f.SetPixel(m.x, m.y, mygifs.Blue)
		}
	}

	f = g.AddCopyFrame()
	maxDist := 0
	maxMark := struct{ x, y int }{}
	for _, m := range aMarkers {
		dist := abs(m.x) + abs(m.y)
		if dist > maxDist {
			maxDist = dist
			maxMark = m
		}
	}
	for _, m := range bMarkers {
		dist := abs(m.x) + abs(m.y)
		if dist > maxDist {
			maxDist = dist
		}
	}
	f.DrawMarker(maxMark.x, maxMark.y, 5, mygifs.Green)
	fmt.Println(maxDist)

	maxSplit := 0
	maxSplitA, maxSplitB := struct{ x, y int }{}, struct{ x, y int }{}
	for _, ma := range aMarkers {
		for _, mb := range bMarkers {
			dist := abs(ma.x-mb.x) + abs(ma.y-mb.y)
			if dist > maxSplit {
				maxSplit = dist
				maxSplitA = ma
				maxSplitB = mb
			}
		}
	}
	fmt.Println(maxSplit)
	f.DrawLine(maxSplitA.x, maxSplitA.y, maxSplitB.x, maxSplitB.y, mygifs.Cyan)
	g.FreezeFrame(100)
	g.Write("solution.gif")
}
