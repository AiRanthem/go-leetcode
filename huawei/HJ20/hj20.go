package main

import "fmt"

// 密码验证
func main() {
	var s string
	for {
		if _, err := fmt.Scanln(&s); err == nil {
			if validate(s) {
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		} else {
			break
		}
	}
}

func validate(s string) bool {
	if len(s) <= 8 {
		return false
	}
	var upper, lower, number, other int
	prefix := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			upper = 1
		} else if c >= 'a' && c <= 'z' {
			lower = 1
		} else if c >= '0' && c <= '9' {
			number = 1
		} else {
			other = 1
		}
		for _, j := range prefix[c] {
			for k := 0; k < len(s)-i; k++ {
				if s[i+k] != s[j+k] {
					break
				}
				if k > 1 {
					return false
				}
			}
		}
		prefix[c] = append(prefix[c], i)
	}
	if upper+lower+number+other < 3 {
		return false
	}
	return true
}
