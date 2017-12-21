package dtemplate

import (
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
)

func init() {
	registry.RegisterDay(template, main)
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("dtemplate/input.txt")

	return "", ""
}
