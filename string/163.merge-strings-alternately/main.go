package main

// https://leetcode.com/problems/merge-strings-alternately/
func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	l := max(l1, l2)
	res := []byte{}

	for i := 0; i < l; i++ {
		if i < l1 {
			res = append(res, word1[i])
		}

		if i < l2 {
			res = append(res, word2[i])
		}
	}

	return string(res)
}
