package main

import "fmt"

const startA = 883
const startB = 879
const factorA = 16807
const factorB = 48271
const div = 2147483647

func main() {
	judged := 0
	a := startA
	b := startB
	for i := 0; i < 40000000; i++ {
		a = (a * factorA) % div
		b = (b * factorB) % div
		if a&0xffff == b&0xffff {
			judged++
		}
	}
	fmt.Println(judged)

	judged = 0
	a = startA
	b = startB
	for i := 0; i < 5000000; i++ {
		a = (a * factorA) % div
		for a%4 != 0 {
			a = (a * factorA) % div
		}
		b = (b * factorB) % div
		for b%8 != 0 {
			b = (b * factorB) % div
		}
		if a&0xffff == b&0xffff {
			judged++
		}
	}
	fmt.Println(judged)
}
