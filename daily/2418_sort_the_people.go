package daily

import "sort"

type Person struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	p := make([]Person, len(names))
	for i := 0; i < len(names); i++ {
		p[i] = Person{names[i], heights[i]}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].Height > p[j].Height
	})
	for i := 0; i < len(names); i++ {
		names[i] = p[i].Name
	}
	return names
}
