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

import (
	"strconv"
)

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

// leetcode 10: https://leetcode.com/problems/regular-expression-matching/
func RegularIsMatch(s string, p string) bool {
	memo := make(map[string]bool)

	var regularIsMatch func(string, int, string, int) bool
	regularIsMatch = func(s string, index int, p string, pindex int) bool {
		if len(p) == pindex {
			return len(s) == index
		}
		if len(s) == index {
			if (len(p)-pindex)%2 == 1 {
				return false
			}
			for i := pindex; i+1 < len(p); i = i + 2 {
				if p[i+1] != '*' {
					return false
				}
			}
			return true
		}

		key := strconv.Itoa(index) + "_" + strconv.Itoa(pindex)
		if value, exist := memo[key]; exist {
			return value
		}

		var match bool
		if s[index] == p[pindex] || p[pindex] == '.' {
			if pindex+1 < len(p) && p[pindex+1] == '*' {
				match = regularIsMatch(s, index, p, pindex+2) || // 0次
					regularIsMatch(s, index+1, p, pindex) // 1次 或 多次
			} else {
				match = regularIsMatch(s, index+1, p, pindex+1)
			}
		} else {
			if pindex+1 < len(p) && p[pindex+1] == '*' {
				match = regularIsMatch(s, index, p, pindex+2)
			} else {
				match = false
			}
		}

		memo[key] = match
		return match
	}

	return regularIsMatch(s, 0, p, 0)
}

// leetcode 28: https://leetcode.com/problems/implement-strstr/
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for index := 0; index < len(haystack); index++ {
		var nindex int
		for index+nindex < len(haystack) && nindex < len(needle) &&
			haystack[index+nindex] == needle[nindex] {
			nindex = nindex + 1
		}

		if nindex == len(needle) {
			return index
		}
	}
	return -1
}
