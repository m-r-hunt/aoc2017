package d4

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strings"
)

func init() {
	registry.RegisterDay(4, main)
}

func validPassphrase(l string) bool {
	words := strings.Split(l, " ")

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
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
			letterCounts[i][c-'a']++
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if letterCounts[i] == letterCounts[j] {
				return false
			}
		}
	}
	return true
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d4/input.txt")
	count, count2 := 0, 0
	for _, l := range lines {
		if validPassphrase(l) {
			count++
		}
		if validPassphrase2(l) {
			count2++
		}
	}
	return fmt.Sprint(count), fmt.Sprint(count2)
}
