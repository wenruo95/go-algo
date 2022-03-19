/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : leetcode430.go
*   coder: zemanzeng
*   date : 2022-03-19 22:08:11
*   desc : leetcode 430~439
*
================================================================*/

package logic

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
