package leetcode

import "fmt"

func getFolderNames(names []string) []string {
	counter := make(map[string]int)
	for i, name := range names {
		nname := name
		n, ok := counter[nname]
		if ok {
			for mapContains(counter, nname) {
				n++
				nname = fmt.Sprintf("%s(%d)", name, n)
			}
		}
		names[i] = nname
		counter[nname] = 0
		counter[name] = n
	}
	return names
}

func mapContains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
