package main

// https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) (res int) {
	counter := map[uint8]int{}

	for left, right := 0, 0; right < len(s); right++ {
		counter[s[right]]++

		for (right-left+1)-maxValue(counter) > k {
			counter[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return
}

func maxValue(m map[uint8]int) (res int) {
	for _, v := range m {
		res = max(res, v)
	}

	return
}
