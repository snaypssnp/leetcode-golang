package main

// https://leetcode.com/problems/permutation-in-string/
func checkInclusion(s2 string, s1 string) bool {
	var n1, n2 = len(s1), len(s2)

	if n1 < n2 {
		return false
	}

	var arr1, arr2 = [26]int{}, [26]int{}

	for _, char := range s2 {
		arr2[char-'a']++
	}

	for i, char := range s1 {
		arr1[char-'a']++
		if i >= n2 {
			arr1[s1[i-n2]-'a']--
		}

		if arr1 == arr2 {
			return true
		}
	}

	return false
}
