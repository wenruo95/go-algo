/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : misc.go
*   coder: zemanzeng
*   date : 2022-03-13 18:39:06
*   desc : 测试用例的一些公共函数
*
================================================================*/

package logic

import "sort"

func stringListEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func intListEqual(dst, src []int) bool {
	if len(dst) != len(src) {
		return false
	}
	for i := 0; i < len(dst); i++ {
		if dst[i] != src[i] {
			return false
		}
	}
	return true
}

func stringListItemEqual(dst, src []string) bool {
	if len(dst) != len(src) {
		return false
	}

	sort.Strings(dst)
	sort.Strings(src)
	for i := 0; i < len(dst); i++ {
		if dst[i] != src[i] {
			return false
		}
	}

	return true
}
