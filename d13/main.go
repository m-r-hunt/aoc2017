package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func main() {
	lines := mygifs.JustLoadLines("input.txt")
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
	fmt.Println(firewalls)
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
				}
			}
		}
		if !caught {
			fmt.Println(delay)
			break
		}
		//fmt.Println(severity)
	}
}
