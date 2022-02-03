/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic_test.go
*   coder: zemanzeng
*   date : 2022-02-03 01:16:43
*   desc :
*
================================================================*/

package logic

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	type parenthesesResult struct {
		s     string
		valid bool
	}
	results := []*parenthesesResult{
		{
			s:     "()",
			valid: true,
		},
		{
			s:     "()[]{}",
			valid: true,
		},
		{
			s:     "(]",
			valid: false,
		},
	}

	for _, result := range results {
		if valid := IsValidParentheses(result.s); valid != result.valid {
			t.Errorf("is_valid_parentheses result:%+v valid:%v", result, valid)
		}
	}

}
