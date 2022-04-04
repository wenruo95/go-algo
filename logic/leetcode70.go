/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode70.go
*   coder: zemanzeng
*   date : 2022-04-03 22:01:46
*   desc : leetcode 70~79
*
================================================================*/

package logic

import (
	"strings"
)

// leetcode 70: https://leetcode.com/problems/climbing-stairs/
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	n1, n2, x := 1, 2, 2
	for x < n {
		x = x + 1
		n1, n2 = n2, n1+n2
	}
	return n2
}

// leetcode 71: https://leetcode.com/problems/simplify-path/
func SimplifyPath(path string) string {
	arr, start := make([]string, 0), 0
	for i := 0; i < len(path); i++ {
		var item string
		if path[i] == '/' || i == len(path)-1 {
			if path[i] == '/' {
				item = path[start:i]
				start = i + 1
			} else {
				item = path[start : i+1]
			}
		}
		if len(item) == 0 {
			continue
		}

		switch item {
		case ".":
			continue

		case "..":
			for j := len(arr) - 1; j >= 0; j-- {
				if len(arr[j]) > 0 {
					arr[j] = ""
					break
				}
			}

		default:
			arr = append(arr, item)
		}

	}

	s := &strings.Builder{}
	s.Grow(len(path))
	for _, item := range arr {
		if len(item) == 0 {
			continue
		}
		s.WriteString("/" + item)
	}

	if s.Len() == 0 {
		return "/"
	}
	return s.String()
}
