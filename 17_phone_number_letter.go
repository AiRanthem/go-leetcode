package leetcode

var phoneNumberMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return phoneNumberMap[digits[0]]
	}
	combs := letterCombinations(digits[1:])
	letters := phoneNumberMap[digits[0]]
	var results []string
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(combs); j++ {
			results = append(results, letters[i]+combs[j])
		}
	}
	return results
}
