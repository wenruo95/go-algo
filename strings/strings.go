/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : strings.go
*   coder: zemanzeng
*   date : 2022-02-02 14:14:46
*   desc : strings
*
================================================================*/

package strings

// leetcode 438: https://leetcode.com/problems/find-all-anagrams-in-a-string/
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	var sIndexs, pIndexs [26]int
	for _, c := range p {
		pIndexs[c-'a']++
	}
	for i := 0; i < len(p); i++ {
		sIndexs[s[i]-'a']++
	}

	indexs := make([]int, 0)
	if pIndexs == sIndexs {
		indexs = append(indexs, 0)
	}

	for index := len(p); index < len(s); index++ {
		sIndexs[s[index-len(p)]-'a']--
		sIndexs[s[index]-'a']++

		if sIndexs == pIndexs {
			indexs = append(indexs, index-len(p)+1)
		}
	}

	return indexs
}

// leetcode 17: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func LetterCombinations(digits string) []string {
	digitStr := [10][]string{
		0: {},
		1: {},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	var indexs []int
	for _, digit := range digits {
		indexs = append(indexs, int(digit-'0'))
	}

	var fn func(indexs []int, cur int) []string
	fn = func(indexs []int, cur int) []string {
		if cur == len(indexs) {
			return nil
		}

		strs := digitStr[indexs[cur]]
		afStrs := fn(indexs, cur+1)
		if len(afStrs) == 0 {
			return strs
		}

		list := make([]string, 0)
		for _, str := range strs {
			for _, afStr := range afStrs {
				list = append(list, str+afStr)
			}
		}

		return list
	}

	return fn(indexs, 0)
}
