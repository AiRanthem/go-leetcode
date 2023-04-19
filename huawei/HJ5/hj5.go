package main

import "fmt"

// 进制转换
func main() {
	var hex string
	_, _ = fmt.Scan(&hex)
	base := 1
	charMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}
	var result int
	for i := len(hex) - 1; i > 1; i-- {
		result += base * charMap[hex[i]]
		base *= 16
	}
	fmt.Println(result)
}
