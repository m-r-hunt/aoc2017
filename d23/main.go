package dtemplate

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"math"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(23, main)
}

type opcode int

const (
	set opcode = iota
	sub
	mul
	jnz
)

type arg struct {
	value      int
	isRegister bool
}

func (a arg) get(registers []int) *int {
	v := &a.value
	if a.isRegister {
		v = &registers[*v]
	}
	return v
}

func (a *arg) parse(s string) {
	if s[0] >= 'a' && s[0] <= 'z' {
		a.value = int(s[0] - 'a')
		a.isRegister = true
	} else {
		a.value, _ = strconv.Atoi(s)
		a.isRegister = false
	}
}

type instruction struct {
	op   opcode
	arg1 arg
	arg2 arg
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d23/input.txt")
	instructions := make([]instruction, len(lines))
	for i, l := range lines {
		f := strings.Fields(l)
		switch f[0] {
		case "set":
			instructions[i].op = set
		case "sub":
			instructions[i].op = sub
		case "mul":
			instructions[i].op = mul
		case "jnz":
			instructions[i].op = jnz
		default:
			panic("Bad instruction")
		}
		instructions[i].arg1.parse(f[1])
		if len(f) >= 3 {
			instructions[i].arg2.parse(f[2])
		}
	}
	pc := 0
	registers := make([]int, 26)
	muls := 0
	for pc >= 0 && pc < len(instructions) {
		i := instructions[pc]
		switch i.op {
		case set:
			*i.arg1.get(registers) = *i.arg2.get(registers)
		case sub:
			*i.arg1.get(registers) -= *i.arg2.get(registers)
		case mul:
			*i.arg1.get(registers) *= *i.arg2.get(registers)
			muls++
		case jnz:
			if *i.arg1.get(registers) != 0 {
				pc += *i.arg2.get(registers)
				continue
			}
		default:
			panic("Instruction not implemented")
		}
		pc++
	}

	h := 0
loop:
	for b := 109300; b <= 126300; b += 17 {
		if b%2 == 0 {
			h++
			continue loop
		}
		max := int(math.Floor(math.Sqrt(float64(b))))
		for n := 3; n <= max; n += 2 {
			if b%n == 0 {
				h++
				continue loop
			}
		}
	}

	return fmt.Sprint(muls), fmt.Sprint(h)
}
