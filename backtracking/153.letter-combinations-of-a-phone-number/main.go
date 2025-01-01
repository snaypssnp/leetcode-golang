package main

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var m = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) (res []string) {
	l := len(digits)

	if l == 0 {
		return
	}

	var backtracing func(curr []rune, i int)

	backtracing = func(curr []rune, i int) {
		if i > l {
			return
		}

		if len(curr) == l {
			res = append(res, string(curr))

			return
		}

		letters := m[rune(digits[i])]

		for _, char := range letters {
			curr = append(curr, char)
			backtracing(curr, i+1)
			curr = curr[:len(curr)-1]
		}
	}

	backtracing([]rune{}, 0)
	return
}
