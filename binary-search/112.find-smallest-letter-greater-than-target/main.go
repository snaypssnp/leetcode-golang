package main

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
func nextGreatestLetter(letters []byte, target byte) (ans byte) {
	l := len(letters) - 1
	ans = letters[0]

	for left, right := 0, l; left <= right; {
		mid := left + (right-left)/2
		letter := letters[mid]

		if letter <= target {
			left = mid + 1
		} else {
			right = mid - 1

			ans = byte(letter)
		}
	}

	return
}
