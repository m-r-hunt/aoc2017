package d9

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(9, main)
}

func main() (string, string) {
	l := mygifs.JustLoadLines("d9/input.txt")[0]
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

	return fmt.Sprint(totalScore), fmt.Sprint(garboChars)
}
