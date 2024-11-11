package main

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) (res []int) {
	if len(p) > len(s) {
		return
	}

	np := len(p)
	ns := len(s)

	mapP := map[uint8]int{}
	for i := 0; i < np; i++ {
		mapP[p[i]]++
	}

	mapS := map[uint8]int{}
	for i := 0; i < np-1; i++ {
		mapS[s[i]]++
	}

core:
	for end := np - 1; end < ns; end++ {
		start := end - (np - 1)

		if start > 0 {
			char := s[start-1]
			if mapS[char] == 1 {
				delete(mapS, char)
			} else {
				mapS[char]--
			}
		}

		mapS[s[end]]++

		for char, v1 := range mapS {
			if v2, ok := mapP[char]; !ok || v1 != v2 {
				continue core
			}
		}

		res = append(res, start)
	}

	return
}