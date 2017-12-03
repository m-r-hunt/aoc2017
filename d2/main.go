package main

import (
	"github.com/m-r-hunt/mygifs"
	"strings"
	"fmt"
	"strconv"
)

const colwidth = 50
const rowheight = 20

func main() {
	lines := mygifs.JustLoadLines("input.txt")
	g := mygifs.NewGif(850, 360)
	defer g.Write("solution.gif")
	checksum := 0
	minis := make([]int, len(lines))
	maxis := make([]int, len(lines))
	for i, l := range lines {
		nums := strings.Split(l, "\t")
		min, max := 100000, 0
		for l, n := range nums {
			nn, _ := strconv.Atoi(n)
			if nn < min {
				min = nn
				minis[i] = l
			}
			if nn > max {
				max = nn
				maxis[i] = l
			}
		}
		checksum += max - min
	}
	fmt.Println(checksum)

	f := g.AddBlankFrame()
	f.DrawText(0, 0, "Checksum", mygifs.Black)
	for i,l := range lines {
		nums := strings.Split(l, "\t")
		for l, n := range nums {
			nn, _ := strconv.Atoi(n)
			c := mygifs.Black
			if l == minis[i] {
				c = mygifs.Red
			} else if l == maxis[i] {
				c = mygifs.Green
			}
			f.DrawText(l * colwidth, (i+1) * rowheight, strconv.Itoa(nn), c)
		}
		max, _ := strconv.Atoi(nums[maxis[i]])
		min, _ := strconv.Atoi(nums[minis[i]])
		f.DrawText(len(nums) * colwidth, (i+1) * rowheight, strconv.Itoa(max - min), mygifs.Blue)
	}
	f.DrawText(len(strings.Split(lines[0], "\t")) * colwidth, (len(lines)+1) * rowheight, strconv.Itoa(checksum), mygifs.Cyan)
	g.FreezeFrame(200)

	topis := make([]int, len(lines))
	botis := make([]int, len(lines))
	evenSum := 0
	for i, l := range lines {
		nums := strings.Split(l, "\t")
		for k, n := range nums {
			for l, m := range nums {
				nn, _ := strconv.Atoi(n)
				mm, _ := strconv.Atoi(m)
				if nn > mm && nn % mm == 0 {
					evenSum += nn / mm
					topis[i] = k
					botis[i] = l
				}
			}
		}
	}
	fmt.Println(evenSum)

	f = g.AddBlankFrame()
	f.DrawText(0, 0, "Even Values", mygifs.Black)
	for i,l := range lines {
		nums := strings.Split(l, "\t")
		for l, n := range nums {
			nn, _ := strconv.Atoi(n)
			c := mygifs.Black
			if l == topis[i] {
				c = mygifs.Red
			} else if l == botis[i] {
				c = mygifs.Green
			}
			f.DrawText(l * colwidth, (i+1) * rowheight, strconv.Itoa(nn), c)
		}
		top, _ := strconv.Atoi(nums[topis[i]])
		bot, _ := strconv.Atoi(nums[botis[i]])
		f.DrawText(len(nums) * colwidth, (i+1) * rowheight, strconv.Itoa(top / bot), mygifs.Blue)
	}
	f.DrawText(len(strings.Split(lines[0], "\t")) * colwidth, (len(lines)+1) * rowheight, strconv.Itoa(evenSum), mygifs.Cyan)
	g.FreezeFrame(200)
}
