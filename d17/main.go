package d17

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(17, main)
}

const steps = 386

type list struct {
	next  *list
	value int
}

func bruteForce() {
	clist := &list{nil, 0}
	clist.next = clist
	for i := 1; i <= 50000000; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		for n := 0; n < steps; n++ {
			clist = clist.next
		}
		clist.next = &list{clist.next, i}
		clist = clist.next
	}
	for {
		if clist.value == 0 {
			fmt.Println(clist.next.value)
			break
		}
		clist = clist.next
	}

}

func main() (string, string) {
	clist := &list{nil, 0}
	clist.next = clist
	ans1 := 0
	for i := 1; i <= 2017; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		for n := 0; n < steps; n++ {
			clist = clist.next
		}
		clist.next = &list{clist.next, i}
		clist = clist.next
	}
	for {
		if clist.value == 2017 {
			ans1 = clist.next.value
			break
		}
		clist = clist.next
	}

	afterZero := 0
	pos := 0
	len := 1
	for i := 1; i <= 50000000; i++ {
		pos = (pos + steps) % len
		if pos == 0 {
			afterZero = i
		}
		len++
		pos++
	}
	return fmt.Sprint(ans1), fmt.Sprint(afterZero)
}
