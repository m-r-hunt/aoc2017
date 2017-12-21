package main

import "fmt"

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

func main() {
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
	fmt.Println(afterZero)
}
