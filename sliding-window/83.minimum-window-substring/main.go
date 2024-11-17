package main

import (
	"math"
)

//https://leetcode.com/problems/minimum-window-substring/

// my solution
func minWindow1(s string, t string) string {
	var res string

	for i := 0; i < len(s); i++ {
		var m = map[uint8]int{}

		for x := 0; x < len(t); x++ {
			char := t[x]
			m[char]++
		}

		for j := i; j < len(s); j++ {
			char := s[j]

			if _, ok := m[char]; ok {
				m[char]--

				if m[char] == 0 {
					delete(m, char)
				}

			}

			if len(m) == 0 {
				if len(res) == 0 || len(res) > (j-i+1) {
					res = s[i : j+1]
					break
				}
			}
		}
	}

	return res
}

func minWindow2(s string, t string) string {
	if t == "" {
		return ""
	}
	var counterT = map[uint8]int{}
	var counterS = map[uint8]int{}
	var res, resLen = []int{-1, -1}, math.MaxInt32

	for i := 0; i < len(t); i++ {
		counterT[t[i]]++
		counterS[t[i]] = 0
	}

	var need, have = 0, len(counterT)

	var left int

	for right := 0; right < len(s); right++ {
		var char = s[right]

		if _, ok := counterS[char]; ok {
			counterS[char]++

			if counterS[char] == counterT[char] {
				need++
			}

		}

		for need == have {
			var newLen = right - left + 1

			if resLen > newLen {
				res = []int{left, right}
				resLen = newLen
			}

			if _, ok := counterT[s[left]]; ok {
				counterS[s[left]]--

				if counterS[s[left]] < counterT[s[left]] {
					need--
				}
			}

			left++
		}
	}

	if resLen == math.MaxInt32 {
		return ""
	}

	return s[res[0] : res[1]+1]
}
