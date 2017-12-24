package d24

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(24, main)
}

type component struct {
	ports [2]int
}

type state struct {
	usedComponents map[int]bool
	nextPort       int
	total          int
	length         int
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d24/input.txt")
	components := make([]component, len(lines))
	for i, l := range lines {
		ps := strings.Split(l, "/")
		components[i].ports[0], _ = strconv.Atoi(ps[0])
		components[i].ports[1], _ = strconv.Atoi(ps[1])
	}

	max := 0
	maxLen := 0
	maxLenStr := 0
	toCheck := []state{state{map[int]bool{}, 0, 0, 0}}
	for len(toCheck) > 0 {
		s := toCheck[0]
		toCheck = toCheck[1:]
		if s.total > max {
			max = s.total
		}
		if s.length > maxLen || (s.length == maxLen && s.total > maxLenStr) {
			maxLen = s.length
			maxLenStr = s.total
		}
		for i := range components {
			if components[i].ports[0] == s.nextPort && !s.usedComponents[i] {
				next := state{}
				next.nextPort = components[i].ports[1]
				next.total = s.total + components[i].ports[0] + components[i].ports[1]
				next.usedComponents = map[int]bool{}
				for k, v := range s.usedComponents {
					next.usedComponents[k] = v
				}
				next.usedComponents[i] = true
				next.length = s.length + 1
				toCheck = append(toCheck, next)
			} else if components[i].ports[1] == s.nextPort && !s.usedComponents[i] {
				next := state{}
				next.nextPort = components[i].ports[0]
				next.total = s.total + components[i].ports[0] + components[i].ports[1]
				next.usedComponents = map[int]bool{}
				for k, v := range s.usedComponents {
					next.usedComponents[k] = v
				}
				next.usedComponents[i] = true
				next.length = s.length + 1
				toCheck = append(toCheck, next)
			}
		}
	}

	return fmt.Sprint(max), fmt.Sprint(maxLenStr)
}
