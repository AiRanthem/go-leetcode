package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 坐标移动
func main() {
	x, y := 0, 0
	var s string
	_, _ = fmt.Scan(&s)
	inputs := strings.Split(s, ";")
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		if len(input) == 0 {
			continue
		}
		if n, err := strconv.Atoi(input[1:]); err == nil && n > 0 {
			switch input[0] {
			case 'W':
				y += n
			case 'S':
				y -= n
			case 'A':
				x -= n
			case 'D':
				x += n
			}
		}
	}
	fmt.Printf("%d,%d", x, y)
}
