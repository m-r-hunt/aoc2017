package d13

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(13, main)
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d13/input.txt")
	firewalls := map[int]int{}
	layerMax := 0
	for _, l := range lines {
		f := strings.Split(l, ": ")
		layer, _ := strconv.Atoi(f[0])
		depth, _ := strconv.Atoi(f[1])
		firewalls[layer] = depth
		if layer > layerMax {
			layerMax = layer
		}
	}

	zeroDelaySeverity := 0
	delay := -1
	for {
		delay++
		ps := delay - 1
		severity := 0
		caught := false
		for ps-delay <= layerMax {
			ps++
			if depth, ok := firewalls[ps-delay]; ok {
				period := (depth - 1) * 2
				if ((ps) % period) == 0 {
					severity += ps * depth
					caught = true
					if delay != 0 {
						break
					}
				}
			}
		}
		if delay == 0 {
			zeroDelaySeverity = severity
		}
		if !caught {
			break
		}
	}

	return fmt.Sprint(zeroDelaySeverity), fmt.Sprint(delay)
}
