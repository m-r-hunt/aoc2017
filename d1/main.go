package d1

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
)

func init() {
	registry.RegisterDay(1, main)
}

func main() (string, string) {
	s := mygifs.JustLoadLines("d1/input.txt")[0]

	nextTotal := 0
	halfTotal := 0
	for i, d := range s {
		digitVal := int(d - '0')
		next := (i + 1) % len(s)
		if s[i] == s[next] {
			nextTotal += digitVal
		}
		half := (i + len(s)/2) % len(s)
		if s[i] == s[half] {
			halfTotal += digitVal
		}
	}
	return fmt.Sprint(nextTotal), fmt.Sprint(halfTotal)
}
