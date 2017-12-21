package d20

import (
	"fmt"
	"github.com/m-r-hunt/aoc2017/registry"
	"github.com/m-r-hunt/mygifs"
	"regexp"
	"strconv"
)

func init() {
	registry.RegisterDay(20, main)
}

type coords struct {
	x, y, z int
}

type particle struct {
	pos, vel, acc coords
	dead          bool
}

func add(a, b coords) coords {
	return coords{a.x + b.x, a.y + b.y, a.z + b.z}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() (string, string) {
	lines := mygifs.JustLoadLines("d20/input.txt")
	particles := make([]particle, len(lines))
	re := regexp.MustCompile("p=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, v=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, a=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>")
	for i, l := range lines {
		m := re.FindStringSubmatch(l)
		particles[i].pos.x, _ = strconv.Atoi(m[1])
		particles[i].pos.y, _ = strconv.Atoi(m[2])
		particles[i].pos.z, _ = strconv.Atoi(m[3])

		particles[i].vel.x, _ = strconv.Atoi(m[4])
		particles[i].vel.y, _ = strconv.Atoi(m[5])
		particles[i].vel.z, _ = strconv.Atoi(m[6])

		particles[i].acc.x, _ = strconv.Atoi(m[7])
		particles[i].acc.y, _ = strconv.Atoi(m[8])
		particles[i].acc.z, _ = strconv.Atoi(m[9])
	}

	for i := 0; i < 1000; i++ {
		for i := range particles {
			particles[i].vel = add(particles[i].vel, particles[i].acc)
			particles[i].pos = add(particles[i].pos, particles[i].vel)
		}

		for i := range particles {
			if !particles[i].dead {
				for j := i + 1; j < len(particles); j++ {
					if !particles[j].dead && particles[i].pos == particles[j].pos {
						particles[i].dead = true
						particles[j].dead = true
					}
				}
			}
		}
	}

	mindist := 10000000
	mini := -1
	for i := range particles {
		dist := abs(particles[i].pos.x) + abs(particles[i].pos.y) + abs(particles[i].pos.z)
		if dist < mindist {
			mini = i
			mindist = dist
		}
	}
	count := 0
	for i := range particles {
		if !particles[i].dead {
			count++
		}
	}
	return fmt.Sprint(mini), fmt.Sprint(count)
}
