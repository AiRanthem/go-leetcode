package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	counter := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		counter[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if n := counter[ransomNote[i]]; n <= 0 {
			return false
		} else {
			counter[ransomNote[i]] = n - 1
		}
	}
	return true
}
