package d12

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(12, main)
}

type dude struct {
	visited     bool
	connections []int
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d12/input.txt")
	dudes := make([]dude, len(lines))
	for i, l := range lines {
		f := strings.Fields(l)
		c := []int{}
		for j := 2; j < len(f); j++ {
			if f[j][len(f[j])-1] == ',' {
				f[j] = f[j][0 : len(f[j])-1]
			}
			n, _ := strconv.Atoi(f[j])
			c = append(c, n)
		}
		dudes[i] = dude{false, c}
	}

	toVisit := []int{0}
	count := 0
	for len(toVisit) > 0 {
		count++
		tv := toVisit[0]
		toVisit = toVisit[1:]
		dudes[tv].visited = true
		for _, v := range dudes[tv].connections {
			if !dudes[v].visited {
				toVisit = append(toVisit, v)
			}
		}
	}

	groupCount := 1
	for i := range dudes {
		if !dudes[i].visited {
			groupCount++
			toVisit := []int{i}
			for len(toVisit) > 0 {
				tv := toVisit[0]
				toVisit = toVisit[1:]
				dudes[tv].visited = true
				for _, v := range dudes[tv].connections {
					if !dudes[v].visited {
						toVisit = append(toVisit, v)
					}
				}
			}

		}
	}
	return fmt.Sprint(count), fmt.Sprint(groupCount)
}
