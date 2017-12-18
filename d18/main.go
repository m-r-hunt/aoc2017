package main

import (
	"fmt"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
	"time"
)

type opcode int

const (
	snd opcode = iota
	set
	add
	mul
	mod
	rcv
	jgz
)

type instruction struct {
	op           opcode
	arg1         int
	arg1register bool
	arg2         int
	arg2register bool
}

func main() {

	// Parse instructions
	lines := mygifs.JustLoadLines("input.txt")
	instructions := make([]instruction, len(lines))
	for i, l := range lines {
		f := strings.Fields(l)
		switch f[0] {
		case "snd":
			instructions[i].op = snd
		case "set":
			instructions[i].op = set
		case "add":
			instructions[i].op = add
		case "mul":
			instructions[i].op = mul
		case "mod":
			instructions[i].op = mod
		case "rcv":
			instructions[i].op = rcv
		case "jgz":
			instructions[i].op = jgz
		default:
			panic("Bad instruction")
		}
		if f[1][0] >= 'a' && f[1][0] <= 'z' {
			instructions[i].arg1 = int(f[1][0] - 'a')
			instructions[i].arg1register = true
		} else {
			instructions[i].arg1, _ = strconv.Atoi(f[1])
			instructions[i].arg1register = false
		}
		if len(f) >= 3 {
			if f[2][0] >= 'a' && f[2][0] <= 'z' {
				instructions[i].arg2 = int(f[2][0] - 'a')
				instructions[i].arg2register = true
			} else {
				instructions[i].arg2, _ = strconv.Atoi(f[2])
				instructions[i].arg2register = false
			}
		}
	}

	c1 := make(chan int, 256)
	c2 := make(chan int, 256)
	sendchans := []chan int{c1, c2}
	rcvchans := []chan int{c2, c1}

	f := func(sndc, rcvc chan int, p int) int {
		pc := 0
		registers := make([]int, 26)
		registers['p'-'a'] = p
		sends := 0
	loop:
		for pc >= 0 && pc < len(instructions) {
			i := instructions[pc]
			switch i.op {
			case snd:
				a := i.arg1
				if i.arg1register {
					a = registers[a]
				}
				sends++
				sndc <- a
			case set:
				a1 := i.arg1
				a2 := i.arg2
				if i.arg2register {
					a2 = registers[a2]
				}
				registers[a1] = a2
			case add:
				a1 := i.arg1
				a2 := i.arg2
				if i.arg2register {
					a2 = registers[a2]
				}
				registers[a1] += a2
			case mul:
				a1 := i.arg1
				a2 := i.arg2
				if i.arg2register {
					a2 = registers[a2]
				}
				registers[a1] *= a2
			case mod:
				a1 := i.arg1
				a2 := i.arg2
				if i.arg2register {
					a2 = registers[a2]
				}
				registers[a1] = registers[a1] % a2
			case rcv:
				select {
				case registers[i.arg1] = <-rcvc:
				case <-time.After(time.Second):
					break loop
				}
			case jgz:
				a1 := i.arg1
				if i.arg1register {
					a1 = registers[a1]
				}
				a2 := i.arg2
				if i.arg2register {
					a2 = registers[a2]
				}
				if a1 > 0 {
					pc += a2
					continue
				}
			default:
				panic("Instruction not implemented")
			}
			pc++
		}
		return sends
	}
	go f(sendchans[0], rcvchans[0], 0)
	out := f(sendchans[1], rcvchans[1], 1)
	fmt.Println(out)
}
