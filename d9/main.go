package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
)

func main() {
	l := mygifs.JustLoadLines("input.txt")[0]
	garbo := false
	depth := 1
	totalScore := 0
	garboChars := 0
	for i := 0; i < len(l); i++ {
		c := l[i]
		if !garbo {
			switch c {
			case '{':
				totalScore += depth
				depth++
			case '}':
				depth--
			case '<':
				garbo = true
			default:
				// Do nothing
			}
		} else {
			switch c {
			case '>':
				garbo = false
			case '!':
				i++
			default:
				garboChars++
			}
		}
	}
	fmt.Println(totalScore)
	fmt.Println(garboChars)
}
