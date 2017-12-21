package d16

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"strconv"
	"strings"
)

func init() {
	registry.RegisterDay(16, main)
}

type opcode int

const (
	s opcode = iota
	x
	p
)

type instruction struct {
	op   opcode
	arg1 int
	arg2 int
}

func main() (string, string) {
	// Parse instructions into a nice form.
	l := mygifs.JustLoadLines("d16/input.txt")[0]
	rawinstrs := strings.Split(l, ",")
	instrs := make([]instruction, len(rawinstrs))
	for i, ri := range rawinstrs {
		switch ri[0] {
		case 's':
			instrs[i].op = s
			instrs[i].arg1, _ = strconv.Atoi(ri[1:])
		case 'x':
			instrs[i].op = x

			ns := strings.Split(ri[1:], "/")
			instrs[i].arg1, _ = strconv.Atoi(ns[0])
			instrs[i].arg2, _ = strconv.Atoi(ns[1])
		case 'p':
			instrs[i].op = p
			instrs[i].arg1 = int(ri[1])
			instrs[i].arg2 = int(ri[3])
		default:
			panic("Unknown instr")
		}
	}

	sxPeriod := 0
	sxPeriodFound := false
	pPeriod := 0
	pPeriodFound := false

	dancers := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	base := 0
	swaps := make([]int, 'p'+1)
	for i := 'a'; i <= 'p'; i++ {
		swaps[int(i)] = int(i)
	}

	ans1 := ""

	for n := 1; true; n++ {
		for _, i := range instrs {
			switch i.op {
			case s:
				base -= i.arg1
				if base < 0 {
					base += 16
				}
			case x:
				dancers[(base+i.arg1)%16], dancers[(base+i.arg2)%16] = dancers[(base+i.arg2)%16], dancers[(base+i.arg1)%16]
			case p:
				swaps[i.arg1], swaps[i.arg2] = swaps[i.arg2], swaps[i.arg1]
			}
		}

		if n == 1 {
			for i := 0; i < 16; i++ {
				d := dancers[(base+i)%16]
				for k, v := range swaps {
					if v == d {
						ans1 += fmt.Sprintf("%c", k)
					}
				}
			}
		}

		if !sxPeriodFound {
			found := true
			for i := range dancers {
				if dancers[(base+i)%16] != 'a'+i {
					found = false
				}
			}
			if found {
				sxPeriod = n
				sxPeriodFound = true
			}
		}
		if !pPeriodFound {
			found := true
			for i := 'a'; i <= 'p'; i++ {
				if swaps[i] != int(i) {
					found = false
				}
			}
			if found {
				pPeriod = n
				pPeriodFound = true
			}
		}
		if sxPeriodFound && pPeriodFound {
			break
		}
	}

	for n := 0; n < 1000000000%sxPeriod; n++ {
		for _, i := range instrs {
			switch i.op {
			case s:
				base -= i.arg1
				if base < 0 {
					base += 16
				}
			case x:
				dancers[(base+i.arg1)%16], dancers[(base+i.arg2)%16] = dancers[(base+i.arg2)%16], dancers[(base+i.arg1)%16]
			}
		}
	}
	for n := 0; n < 1000000000%pPeriod; n++ {
		for _, i := range instrs {
			switch i.op {
			case p:
				swaps[i.arg1], swaps[i.arg2] = swaps[i.arg2], swaps[i.arg1]
			}
		}
	}

	ans2 := ""
	for i := 0; i < 16; i++ {
		d := dancers[(base+i)%16]
		for k, v := range swaps {
			if v == d {
				ans2 += fmt.Sprintf("%c", k)
			}
		}
	}

	return ans1, ans2
}
