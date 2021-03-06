package d18

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
	"time"
)

func init() {
	registry.RegisterDay(18, main)
}

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

func main() (string, string) {

	// Parse instructions
	lines := mygifs.JustLoadLines("d18/input.txt")
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
	c3 := make(chan int, 256)
	sendchans := []chan int{c1, c2, c3}
	rcvchans := []chan int{c2, c1, c3}

	f := func(p int) int {
		pc := 0
		registers := make([]int, 26)
		if p != 2 {
			registers['p'-'a'] = p
		}
		sends := 0
		lastPlayed := 0
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
				lastPlayed = a
				sendchans[p] <- a
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
				if p == 2 {
					a := i.arg1
					if i.arg1register {
						a = registers[a]
					}
					if a != 0 {
						return lastPlayed
					}
				} else {
					select {
					case registers[i.arg1] = <-rcvchans[p]:
					case <-time.After(time.Second):
						break loop
					}
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
	go f(0)
	out := f(1)

	ans1 := f(2)

	return fmt.Sprint(ans1), fmt.Sprint(out)
}
