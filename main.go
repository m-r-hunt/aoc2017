package main

import (
	"github.com/m-r-hunt/mygifs"
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/aoc2017/d1"
)

func main() {
	answers := mygifs.JustLoadLines("answers.txt")
	for day := 1; day <= 25; day++ {
		fmt.Printf("Day %2v ", day)
		part1 := registry.GetPart(day, 1)
		if part1 == nil {
			fmt.Print("[NOT IMPLEMENTED]\n")
			break
		}
		result1 := part1()
		if len(answers) < (day-1)*2+1 {
			fmt.Printf("Part 1 [UNANSWERED] Result: %v", result1)
		} else if result1 != answers[(day-1)*2] {
			fmt.Printf("Part 1 [FAIL] Got: %v, Expected: %v\n", result1, answers[(day-1)*2])
		} else {
			fmt.Print("Part 1 [OK]\n")
		}

		part2 := registry.GetPart(day, 2)
		if part2 == nil {
			fmt.Print("       Part 2 [NOT IMPLEMENTED]\n")
			break
		}
		result2 := part2()
		if len(answers) < (day-1)*2+2 {
			fmt.Printf("       Part 2 [UNANSWERED] Result: %v", result2)
		} else if result1 != answers[(day-1)*2 + 1] {
			fmt.Print("       Part 2 [FAIL] Got: %v, Expected: %v", result2, answers[(day-1)*2 + 1])
		} else {
			fmt.Print("       Part 1 [OK]\n")
		}
	}
}
