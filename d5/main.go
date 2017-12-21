package d5

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
)

func init() {
	registry.RegisterDay(5, main)
}

func interpet(jumps []int, part2 bool) int {
	pc := 0
	steps := 0
	for pc >= 0 && pc < len(jumps) {
		steps++
		npc := pc + jumps[pc]
		if part2 && jumps[pc] >= 3 {
			jumps[pc]--
		} else {
			jumps[pc]++
		}
		pc = npc
	}
	return steps
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d5/input.txt")
	jumps := make([]int, len(lines))
	for i := range lines {
		n, _ := strconv.Atoi(lines[i])
		jumps[i] = n
	}
	jumpsCopy := make([]int, len(lines))
	copy(jumpsCopy, jumps)

	return fmt.Sprint(interpet(jumps, false)), fmt.Sprint(interpet(jumpsCopy, true))
}
