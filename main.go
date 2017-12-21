package main

import (
	"flag"
	"fmt"
	"github.com/dterei/gotsc"
	_ "github.com/m-r-hunt/aoc2017/d1"
	_ "github.com/m-r-hunt/aoc2017/d10"
	_ "github.com/m-r-hunt/aoc2017/d11"
	_ "github.com/m-r-hunt/aoc2017/d12"
	_ "github.com/m-r-hunt/aoc2017/d13"
	_ "github.com/m-r-hunt/aoc2017/d14"
	_ "github.com/m-r-hunt/aoc2017/d15"
	_ "github.com/m-r-hunt/aoc2017/d16"
	_ "github.com/m-r-hunt/aoc2017/d17"
	_ "github.com/m-r-hunt/aoc2017/d18"
	_ "github.com/m-r-hunt/aoc2017/d19"
	_ "github.com/m-r-hunt/aoc2017/d2"
	_ "github.com/m-r-hunt/aoc2017/d20"
	_ "github.com/m-r-hunt/aoc2017/d21"
	_ "github.com/m-r-hunt/aoc2017/d3"
	_ "github.com/m-r-hunt/aoc2017/d4"
	_ "github.com/m-r-hunt/aoc2017/d5"
	_ "github.com/m-r-hunt/aoc2017/d6"
	_ "github.com/m-r-hunt/aoc2017/d7"
	_ "github.com/m-r-hunt/aoc2017/d8"
	_ "github.com/m-r-hunt/aoc2017/d9"
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
	totalTime := time.Duration(0)
	passes, fails := 0, 0
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
		totalTime += time
		fmt.Printf("(%v Cycles)", endtsc-starttsc-tsc)
		fmt.Printf("(%vms):\n", time.Nanoseconds()/1000000)

		if len(answers) < (day-1)*2+1 {
			fmt.Printf("  Part 1 [UNANSWERED] Result: %v\n", result1)
		} else if result1 != answers[(day-1)*2] {
			fails++
			fmt.Printf("  Part 1 [FAIL] Got: %v, Expected: %v\n", result1, answers[(day-1)*2])
		} else {
			passes++
			fmt.Printf("  Part 1 [OK] Result: %v\n", result1)
		}
		if len(answers) < (day-1)*2+2 {
			fmt.Printf("  Part 2 [UNANSWERED] Result: %v\n", result2)
		} else if result2 != answers[(day-1)*2+1] {
			fails++
			fmt.Printf("  Part 2 [FAIL] Got: %v, Expected: %v\n", result2, answers[(day-1)*2+1])
		} else {
			passes++
			fmt.Printf("  Part 2 [OK] Result: %v\n", result2)
		}
	}
	fmt.Printf("Finished. Passed: %v Failed: %v Total Time: %vms", passes, fails, totalTime.Nanoseconds()/1000000)
}
