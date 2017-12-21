package dtemplate

import (
	"github.com/m-r-hunt/mygifs"
	"github.com/m-r-hunt/aoc2017/registry"
)

func init() {
	registry.RegisterDay(template, main)
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("input.txt")

	return "", ""
}
