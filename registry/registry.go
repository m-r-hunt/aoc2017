package registry

type Day func() (string, string)

var registry = map[int]Day{}

func RegisterDay(day int, fn Day) {
	registry[day] = fn
}

func GetDay(day int) Day {
	return registry[day]
}
