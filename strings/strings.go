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

func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	pIndexs := make(map[byte]struct{})
	for _, c := range p {
		pIndexs[byte(c)] = struct{}{}
	}

	indexs := make([]int, 0)
	sIndexs := make(map[byte]int)

	var left int
	for index := 0; index < len(s); index++ {

		// 字符不存在于p, 清除[left, index)
		if _, exist := pIndexs[s[index]]; !exist {
			for j := left; j < index; j++ {
				delete(sIndexs, s[j])
			}
			left = index + 1
			continue
		}

		// 存在于p 且存在于s中,清除[left, oldIndex]
		if oldIndex, exist := sIndexs[s[index]]; exist {
			for j := left; j <= oldIndex; j++ {
				delete(sIndexs, s[j])
			}
			left = oldIndex + 1
		}

		// 然后赋新值
		sIndexs[s[index]] = index

		// 检测判断是否符合预期
		if len(sIndexs) == len(pIndexs) {
			indexs = append(indexs, left)
			continue
		}

	}

	return indexs
}
