package main

import (
	"github.com/m-r-hunt/mygifs"
	"strings"
	"fmt"
)

func validPassphrase(l string) bool {
	words := strings.Split(l, " ")

	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if words[i] == words[j] {
				return false
			}
		}
	}
	return true
}

type letterCount [26]int

func validPassphrase2(l string) bool {
	words := strings.Split(l, " ")
	letterCounts := make([]letterCount, len(words))
	for i, w := range words {
		for _, c := range w {
			letterCounts[i][c - 'a']++
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if letterCounts[i] == letterCounts[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	mygifs.Delay = 2
	g := mygifs.NewGif(250, 30)
	defer g.Write("solution.gif")
	lines := mygifs.JustLoadLines("input.txt")
	count, count2 := 0, 0
	for _, l := range lines {
		f := g.AddBlankFrame()
		f.DrawText(0, 0, l, mygifs.Black)
		if validPassphrase(l) {
			count++
			f.DrawText(0, 15, fmt.Sprintf("Part 1: %3v", count), mygifs.Green)
		} else {
			f.DrawText(0, 15, fmt.Sprintf("Part 1: %3v", count), mygifs.Red)
		}
		if validPassphrase2(l) {
			count2++
			f.DrawText(50, 15, fmt.Sprintf("Part 2: %3v", count2), mygifs.Green)
		} else {
			f.DrawText(50, 15, fmt.Sprintf("Part 2: %3v", count2), mygifs.Red)

		}
	}
	fmt.Println(count)
	fmt.Println(count2)
	g.FreezeFrame(100)
}
