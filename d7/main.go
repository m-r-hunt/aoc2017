package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

type tower struct {
	totalWeight int
	weight      int
	children    []string
}

func weight(root string, towers map[string]tower) int {
	if towers[root].totalWeight != 0 {
		return towers[root].totalWeight
	}
	t := towers[root].weight
	for _, c := range towers[root].children {
		t += weight(c, towers)
	}
	tt := towers[root]
	tt.totalWeight = t
	towers[root] = tt
	return t
}

func findImbalance(root string, towers map[string]tower) bool {
	weights := make([]int, len(towers[root].children))
	for i,c := range towers[root].children {
		weights[i] = weight(c, towers)
	}
	for i := 0; i < len(weights)-1; i++ {
		if weights[i] != weights[i+1] {
			t := false
			if weights[i] != weights[i+2] {
				t= findImbalance(towers[root].children[i], towers)
				if t {
					tt := towers[towers[root].children[i]]
					fmt.Println(tt.weight)
					fmt.Println(tt.totalWeight)
					fmt.Println(towers[towers[root].children[i+1]].totalWeight)
					fmt.Println(towers[towers[root].children[i+1]].weight)
				}
			} else {
				t= findImbalance(towers[root].children[i+1], towers)
				if t {
					tt := towers[towers[root].children[i+1]]
					fmt.Println(tt.weight)
					fmt.Println(tt.totalWeight)
					fmt.Println(towers[towers[root].children[i]].totalWeight)
					fmt.Println(towers[towers[root].children[i]].weight)
				}
			}
			goto found
		}
	}
	return true
	found:

	return false
}

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	towers := map[string]tower{}
	for _, s := range lines {
		fields := strings.Fields(s)
		t := tower{}
		t.children = []string{}
		i, _ := strconv.Atoi(strings.Trim(fields[1], "()"))
		t.weight = i
		if len(fields) >= 4 {
			for i := 3; i < len(fields); i++ {
				s := fields[i]
				if s[len(s)-1] == ',' {
					s = s[0 : len(s)-1]
				}
				t.children = append(t.children, s)
			}
		}
		towers[fields[0]] = t
	}
	fmt.Println(towers)
	dudes := make(map[string]bool)
	for _, ds := range towers {
		for _, s := range ds.children {
			dudes[s] = true
		}
	}
	fmt.Println(dudes)
	root := ""
	for d := range towers {
		if !dudes[d] {
			fmt.Println(d)
			root = d
		}
	}
	findImbalance(root, towers)
}
