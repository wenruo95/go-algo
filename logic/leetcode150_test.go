/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode150_test.go
*   coder: zemanzeng
*   date : 2022-03-19 22:31:58
*   desc :
*
================================================================*/

package logic

import "testing"

func TestReverseWords(t *testing.T) {
	type reverseResult struct {
		input  string
		output string
	}

	results := []*reverseResult{
		{
			input:  "the sky is blue",
			output: "blue is sky the",
		},
		{
			input:  "  hello world  ",
			output: "world hello",
		},
		{
			input:  "a good   example",
			output: "example good a",
		},
	}

	for _, result := range results {
		if output := ReverseWords(result.input); output != result.output {
			t.Errorf("reverse_words result:%+v output[%v]:%v expect[%v]:%v",
				result, len(output), output, len(result.output), result.output)
		}
	}

}
