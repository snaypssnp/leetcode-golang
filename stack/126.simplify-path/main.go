package main

import "strings"

// https://leetcode.com/problems/simplify-path/
func simplifyPath(path string) string {
	stack := []string{}
	items := strings.Split(path, "/")

	for _, item := range items {
		if item == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if item != "." && item != "" {
			stack = append(stack, item)
		}
	}

	return "/" + strings.Join(stack, "/")
}
