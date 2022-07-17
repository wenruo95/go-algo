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
	"encoding/json"
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

func stringArraysEqual(dst, src [][]string, sortFlag bool) bool {
	listA := make([]string, 0, len(dst))
	for _, list := range dst {
		if sortFlag {
			sort.Strings(list)
		}
		listA = append(listA, strings.Join(list, "##"))
	}

	listB := make([]string, 0, len(src))
	for _, list := range src {
		if sortFlag {
			sort.Strings(list)
		}
		listB = append(listB, strings.Join(list, "##"))
	}

	return stringListItemEqual(listA, listB)
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

func arraysEqual(a [][]int, b [][]int, sortFlag bool) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[string]struct{})
	mb := make(map[string]struct{})

	for _, array := range a {
		if sortFlag {
			sort.Ints(array)
		}

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
		if sortFlag {
			sort.Ints(array)
		}

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

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func arrays2str(arrays [][]int) string {
	var s string = "\n"
	for _, array := range arrays {
		for _, item := range array {
			s = s + strconv.Itoa(item) + "\t"
		}
		s = s + "\n"
	}
	return s
}

func minInt(list ...int) int {
	if len(list) == 0 {
		return 0
	}
	var min int = list[0]
	for i := 1; i < len(list); i++ {
		if min > list[i] {
			min = list[i]
		}
	}
	return min
}

func maxInt(list ...int) int {
	if len(list) == 0 {
		return 0
	}
	var max int = list[0]
	for i := 1; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}

func jsonstr(i interface{}) string {
	buff, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(buff)
}

func byteArr2Str(arr [][]byte) string {
	l := make([]string, 0)
	for _, b := range arr {
		l = append(l, string(b))
	}
	return strings.Join(l, ",")
}
