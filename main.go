package main

import (
	"flag"
	"fmt"
	"github.com/dterei/gotsc"
	_ "github.com/m-r-hunt/aoc2017/d1"
	_ "github.com/m-r-hunt/aoc2017/d2"
	_ "github.com/m-r-hunt/aoc2017/d3"
	_ "github.com/m-r-hunt/aoc2017/d4"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"time"
)

func main() {
	tsc := gotsc.TSCOverhead()
	var day = flag.Int("d", -1, "Run specific day")
	flag.Parse()
	if *day != -1 {
		dayFn := registry.GetDay(*day)
		r1, r2 := dayFn()
		fmt.Println(r1)
		fmt.Println(r2)
		return
	}

	answers := mygifs.JustLoadLines("answers.txt")
	for day := 1; day <= 25; day++ {
		fmt.Printf("Day %v ", day)

		dayFn := registry.GetDay(day)
		if dayFn == nil {
			fmt.Print("[NOT IMPLEMENTED]\n")
			break
		}

		start := time.Now()
		starttsc := gotsc.BenchStart()
		result1, result2 := dayFn()
		endtsc := gotsc.BenchEnd()
		time := time.Since(start)
		fmt.Printf("(%v Cycles)", endtsc-starttsc-tsc)
		fmt.Printf("(%vms):\n", time.Nanoseconds()/1000)

		if len(answers) < (day-1)*2+1 {
			fmt.Printf("  Part 1 [UNANSWERED] Result: %v\n", result1)
		} else if result1 != answers[(day-1)*2] {
			fmt.Printf("  Part 1 [FAIL] Got: %v, Expected: %v\n", result1, answers[(day-1)*2])
		} else {
			fmt.Printf("  Part 1 [OK] Result: %v\n", result1)
		}
		if len(answers) < (day-1)*2+2 {
			fmt.Printf("  Part 2 [UNANSWERED] Result: %v\n", result2)
		} else if result2 != answers[(day-1)*2+1] {
			fmt.Print("  Part 2 [FAIL] Got: %v, Expected: %v\n", result2, answers[(day-1)*2+1])
		} else {
			fmt.Printf("  Part 2 [OK] Result: %v\n", result2)
		}
	}
}
