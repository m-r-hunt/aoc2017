package d15

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(15, main)
}

const startA = 883
const startB = 879
const factorA = 16807
const factorB = 48271
const div = 2147483647

func main() (string, string) {
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
	ans1 := judged

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
	return fmt.Sprint(ans1), fmt.Sprint(judged)
}
