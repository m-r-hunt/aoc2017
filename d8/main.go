package d8

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(8, main)
}

type instruction struct {
	target    string
	op        int
	opValue   int
	testReg   string
	test      string
	testValue int
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d8/input.txt")
	instructions := make([]instruction, len(lines))
	allMax := 0
	for i, l := range lines {
		f := strings.Fields(l)
		instructions[i].target = f[0]
		switch f[1] {
		case "inc":
			instructions[i].op = +1
		case "dec":
			instructions[i].op = -1
		default:
			panic(0)
		}
		instructions[i].opValue, _ = strconv.Atoi(f[2])
		instructions[i].testReg = f[4]
		instructions[i].test = f[5]
		instructions[i].testValue, _ = strconv.Atoi(f[6])
	}

	pc := 0
	registers := map[string]int{}
	for pc >= 0 && pc < len(instructions) {
		apply := false
		i := instructions[pc]
		switch i.test {
		case "<":
			if registers[i.testReg] < i.testValue {
				apply = true
			}
		case "<=":
			if registers[i.testReg] <= i.testValue {
				apply = true
			}
		case ">":
			if registers[i.testReg] > i.testValue {
				apply = true
			}
		case ">=":
			if registers[i.testReg] >= i.testValue {
				apply = true
			}
		case "==":
			if registers[i.testReg] == i.testValue {
				apply = true
			}
		case "!=":
			if registers[i.testReg] != i.testValue {
				apply = true
			}
		default:
			panic(i.test)
		}

		if apply {
			registers[i.target] += i.op * i.opValue
		}
		if registers[i.target] > allMax {
			allMax = registers[i.target]
		}
		pc++
	}
	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return fmt.Sprint(max), fmt.Sprint(allMax)
}
