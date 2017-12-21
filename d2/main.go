package d2

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(2, main)
}

const colwidth = 50
const rowheight = 20

func main() (string, string) {
	lines := mygifs.JustLoadLines("d2/input.txt")

	checksum := 0
	for _, l := range lines {
		nums := strings.Split(l, "\t")
		min, max := 100000, 0
		for _, n := range nums {
			nn, _ := strconv.Atoi(n)
			if nn < min {
				min = nn
			}
			if nn > max {
				max = nn
			}
		}
		checksum += max - min
	}

	evenSum := 0
	for _, l := range lines {
		nums := strings.Split(l, "\t")
		for _, n := range nums {
			for _, m := range nums {
				nn, _ := strconv.Atoi(n)
				mm, _ := strconv.Atoi(m)
				if nn > mm && nn%mm == 0 {
					evenSum += nn / mm
				}
			}
		}
	}

	return fmt.Sprint(checksum), fmt.Sprint(evenSum)
}
