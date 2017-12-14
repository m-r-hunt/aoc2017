package main

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/d10"
	"math/bits"
	"strconv"
)

const input = "stpzcrnm"
const test_input = "flqrgnkx"

func main() {
	used := 0
	usedGrid := [128][128]bool{}
	for i := 0; i < 128; i++ {
		kh := d10.KnotHash(input + "-" + strconv.Itoa(i))
		for l, n := range kh {
			used += bits.OnesCount(uint(n))
			for k,j := range []int{7,6,5,4,3,2,1,0} {
				usedGrid[i][l*8+k] = (uint(n) & (1 << uint(j))) != 0
			}
		}
	}
	fmt.Println(used)

	for i := 0 ; i < 128; i++ {
		for j := 0 ; j < 128; j++ {
			if usedGrid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	visited := [128][128]bool{}
	regions := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				if usedGrid[i][j] {
					regions++
					toVisit := []struct{ i, j int }{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
					for len(toVisit) > 0 {
						v := toVisit[0]
						toVisit = toVisit[1:]
						if v.i >= 0 && v.i < 128 && v.j >= 0 && v.j < 128 && !visited[v.i][v.j] {
							visited[v.i][v.j] = true
							if usedGrid[v.i][v.j] {
								toVisit = append(toVisit,
									[]struct{ i, j int }{{v.i - 1, v.j}, {v.i + 1, v.j}, {v.i, v.j - 1}, {v.i, v.j + 1}}...)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(regions)
}
