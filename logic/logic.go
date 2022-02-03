/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic.go
*   coder: zemanzeng
*   date : 2022-02-03 00:49:06
*   desc :
*
================================================================*/

package logic

import "container/list"

// leetcode 20:https://leetcode.com/problems/valid-parentheses/
func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	left := list.New()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			left.PushBack(s[i])
			continue
		}
		if left.Len() != 0 && pairs[s[i]] == left.Back().Value.(byte) {
			left.Remove(left.Back())
			continue
		}
		return false
	}

	return left.Len() == 0
}
