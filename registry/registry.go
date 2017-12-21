package registry

type Part func() string

var registry = map[int][2]Part{}

func RegisterDay(day int, part1, part2 func() string) {
	p := [2]Part{part1, part2}
	registry[day] = p
}

func GetPart(day, part int) Part {
	return registry[day][part-1]
}
