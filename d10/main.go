package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func knotHash(toHash string) string {
	hash := make([]int, 256)
	for i := range hash {
		hash[i] = i
	}

	extraLens := []int{17, 31, 73, 47, 23}
	lens := make([]int, len(toHash), len(toHash)+len(extraLens))
	for i, c := range toHash {
		lens[i] = int(c)
	}
	lens = append(lens, extraLens...)

	pos := 0
	skipSize := 0
	for round := 0; round < 64; round++ {
		for _, length := range lens {
			for i := 0; i < length/2; i++ {
				i1 := (pos + i) % len(hash)
				i2 := (pos + length - 1 - i) % len(hash)
				hash[i1], hash[i2] = hash[i2], hash[i1]
			}
			pos = (pos + length + skipSize) % len(hash)
			skipSize++
		}
	}

	denseHash := make([]int, len(hash)/16)
	for i := range denseHash {
		for j := 0; j < 16; j++ {
			denseHash[i] ^= hash[i*16+j]
		}
	}

	out := ""
	for _, h := range denseHash {
		out += fmt.Sprintf("%02x", h)
	}
	return out
}

func main() {
	line := mygifs.JustLoadLines("input.txt")[0]

	// Part 1: Proto knot hash
	lenstrs := strings.Split(line, ",")
	lens := make([]int, len(lenstrs))
	for i, s := range lenstrs {
		lens[i], _ = strconv.Atoi(s)
	}

	hash := make([]int, 256)
	for i, _ := range hash {
		hash[i] = i
	}

	pos := 0
	skipSize := 0
	for _, length := range lens {
		for i := 0; i < length/2; i++ {
			i1 := (pos + i) % len(hash)
			i2 := (pos + length - 1 - i) % len(hash)
			hash[i1], hash[i2] = hash[i2], hash[i1]
		}
		pos += length + skipSize
		pos = pos % len(hash)
		skipSize++
	}
	fmt.Println(hash[0] * hash[1])

	// Part 2
	fmt.Println(knotHash(line))
}
