package main

import (
	"github.com/m-r-hunt/mygifs"
	"strings"
	"fmt"
)

func validPassphrase(l string) bool {
	words := strings.Split(l, " ")

	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if words[i] == words[j] {
				return false
			}
		}
	}
	return true
}

type letterCount [26]int

func validPassphrase2(l string) bool {
	words := strings.Split(l, " ")
	letterCounts := make([]letterCount, len(words))
	for i, w := range words {
		for _, c := range w {
			letterCounts[i][c - 'a']++
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if letterCounts[i] == letterCounts[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	count := 0
	for _, l := range lines {
		if validPassphrase(l) {
			count++
		}
	}
	fmt.Println(count)

	count2 := 0
	for _, l := range lines {
		if validPassphrase2(l) {
			count2++
		}
	}
	fmt.Println(count2)
}
