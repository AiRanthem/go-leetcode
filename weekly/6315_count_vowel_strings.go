package weekly

func vowelStrings(words []string, left int, right int) (cnt int) {
	var isVowel func(s string) int
	var yuan func(c byte) bool
	yuan = func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}
	isVowel = func(s string) int {
		if yuan(s[0]) && yuan(s[len(s)-1]) {
			return 1
		} else {
			return 0
		}
	}
	for i := left; i <= right; i++ {
		cnt += isVowel(words[i])
	}
	return
}
