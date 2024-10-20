package main

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) (res [][]string) {
	m := map[string][]string{}

	for _, str := range strs {
		key := sortString(str)

		m[key] = append(m[key], str)
	}

	for _, subStrs := range m {
		res = append(res, subStrs)
	}

	return res
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
