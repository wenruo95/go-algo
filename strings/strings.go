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
	"fmt"
	"math"
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

// 字符的全排列问题: 如 "12" 排列有"1" "2" "12"
// TODO
func Permutations(s string) []string {
	chcnt := make(map[string]int)
	for _, ch := range s {
		chcnt[string(ch)] = chcnt[string(ch)] + 1
	}

	var choose func(map[string]int, int) []string
	choose = func(chcnt map[string]int, cnt int) []string {
		if cnt == 1 {
			list := make([]string, 0)
			for ch, cnt := range chcnt {
				if cnt > 0 {
					list = append(list, ch)
				}
			}
			return list
		}

		l1 := make([]string, 0)
		for ch, cnt := range chcnt {
			if cnt <= 0 {
				continue
			}

			chcnt[ch] = cnt - 1
			l2 := choose(chcnt, cnt-1)
			chcnt[ch] = cnt + 1

			for _, item := range l2 {
				l1 = append(l1, ch+item)
			}
		}
		return l1
	}

	l := make([]string, 0)
	for i := 1; i <= len(s); i++ {
		l = append(l, choose(chcnt, i)...)
	}
	return l
}

// string2int
func Str2Int(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	var result int
	var positive bool = true
	for index, item := range s {
		if item == '+' || item == '-' {
			if index != 0 || len(s) == 1 {
				return 0, fmt.Errorf("invalid s[%d]:%s", index, string(item))
			}
			positive = (item != '-')
			continue
		}

		digit := int(item - '0')
		if digit < 0 || digit > 9 {
			return 0, fmt.Errorf("invalid s[%d]:%s", index, string(item))
		}

		if positive {
			if result > math.MaxInt/10 || result*10 > math.MaxInt-digit {
				return 0, fmt.Errorf("s:%v more than:%v", s, math.MaxInt)
			}
		} else {
			if -result < math.MinInt/10 || -result*10 < math.MinInt+digit {
				return 0, fmt.Errorf("s:%v min than:%v", s, math.MinInt)
			}
		}

		result = result*10 + int(digit)
	}

	if positive {
		return result, nil
	}

	return -result, nil
}

// We Are Happy
// We%20Are%20Happy
func ReplaceSpace(s string) string {
	/*
		buff := make([]byte, 0)
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				buff = append(buff, []byte("%20")...)
			} else {
				buff = append(buff, s[i])
			}
		}
		return string(buff)
	*/

	var result string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = result + "%20"
		} else {
			result = result + string(s[i])
		}
	}
	return result
}

// leetcode 30: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func FindSubstring(s string, words []string) []int {
	// 1. create trie tree
	type dataNode struct {
		data  map[byte]*dataNode
		item  byte
		word  string
		count int
	}

	diffCnt := 0
	root := &dataNode{data: make(map[byte]*dataNode)}
	for i := 0; i < len(words); i++ {
		var node *dataNode = root

		for j := 0; j < len(words[i]); j++ {
			b := words[i][j]

			n2, exist := node.data[b]
			if exist {
				node = n2
				continue
			}

			n2 = &dataNode{data: make(map[byte]*dataNode), item: b}
			node.data[b] = n2
			node = n2
		}

		node.word = words[i]
		node.count = node.count + 1
		if node.count == 1 {
			diffCnt = diffCnt + 1
		}
	}

	// 2. 从前往后匹配
	list := make([]int, 0)
	fullLen := len(words) * len(words[0])
	for i := 0; i < len(s); i++ {
		var node *dataNode = root

		matchCnt := 0
		set := make(map[string]int)
		for j := 0; j < fullLen && i+j < len(s); j++ {
			n2, exist := node.data[s[i+j]]
			if !exist {
				break
			}

			if len(n2.word) > 0 { // 匹配到叶子节点
				if set[n2.word] >= n2.count { // 有重复匹配问题
					break
				}
				set[n2.word] = set[n2.word] + 1
				matchCnt = matchCnt + 1

				node = root
				continue
			}

			node = n2
		}

		if len(set) == diffCnt && matchCnt == len(words) {
			list = append(list, i)
			continue
		}
	}

	return list
}
