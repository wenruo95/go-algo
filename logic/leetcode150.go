/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic150.go
*   coder: zemanzeng
*   date : 2022-03-19 22:11:43
*   desc : leetcode 150~159
*
================================================================*/

package logic

// leetcode 151: https://leetcode.com/problems/reverse-words-in-a-string/
func ReverseWords(s string) string {
	edges := make([]int, 0)

	var lastIndex int = -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			lastIndex = i
			continue
		}

		if i+1 >= len(s) || s[i+1] == ' ' {
			edges = append(edges, lastIndex+1, i+1)
		}
	}

	var result string
	for i := len(edges) - 1; i > 0; i = i - 2 {
		left, right := edges[i-1], edges[i]
		if len(result) == 0 {
			result = s[left:right]
			continue
		}
		result = result + " " + s[left:right]
	}
	return result
}
