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

import (
	"sort"
	"strconv"
	"strings"
)

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

func arraysEqual(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[string]struct{})
	mb := make(map[string]struct{})

	for _, array := range a {
		sort.Ints(array)

		var key strings.Builder
		for index, value := range array {
			if index == 0 {
				key.WriteString(strconv.Itoa(value))
			} else {
				key.WriteString("_" + strconv.Itoa(value))
			}
		}
		ma[key.String()] = struct{}{}
	}

	for _, array := range b {
		sort.Ints(array)

		var key strings.Builder
		for index, value := range array {
			if index == 0 {
				key.WriteString(strconv.Itoa(value))
			} else {
				key.WriteString("_" + strconv.Itoa(value))
			}
		}
		mb[key.String()] = struct{}{}

		if _, exist := ma[key.String()]; !exist {
			return false
		}
	}

	return len(ma) == len(mb)
}

// GCD: Greatest common divisor(最大公约数)
func MaxGCD(m, n int) int {
	mod := m % n
	if mod == 0 {
		return n
	}
	return MaxGCD(n, mod)
}

func intMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
