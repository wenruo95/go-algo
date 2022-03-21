/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : logic.go
*   coder: zemanzeng
*   date : 2022-02-03 00:49:06
*   desc : leetcode logic
*
================================================================*/

package logic

import (
	"fmt"
	"math"
	"strconv"
)

// 10=3 + 3 + 4 => ( 3 * 3 * 4 ) = 36
func GetMaxSplitN(n int) int {

	var fn func(n int) int

	memo := make(map[int]int)
	fn = func(n int) int {
		if n <= 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return n
		}
		if maxSplit, exist := memo[n]; exist {
			return maxSplit
		}

		var max int = n
		for i := 1; i < n; i++ {
			if maxSplit := fn(i) * fn(n-i); maxSplit > max {
				max = maxSplit
			}
		}
		memo[n] = max
		return max

	}

	return fn(n)
}

// ababc => abab c
func MaxSplitStringN(s string) []string {
	if len(s) == 0 {
		return nil
	}

	chIndexList := make(map[byte][]int) // first-end
	for i := 0; i < len(s); i++ {
		chIndexList[s[i]] = append(chIndexList[s[i]], i)
	}

	splitStrs := make([]string, 0)

	var index, left, last int = 0, 0, -1
	for index = 0; index < len(s); index++ {
		if index == last {
			splitStrs = append(splitStrs, s[left:last+1])

			left = last + 1
			continue
		}

		if lastIndexs, exist := chIndexList[s[index]]; exist && len(lastIndexs) > 1 {
			lastIndex := lastIndexs[len(lastIndexs)-1]
			if lastIndex > last {
				last = lastIndex
			}
			continue
		}

		if index > last {
			if last == 0 || last < left {
				last = left
			}
			splitStrs = append(splitStrs, s[left:last+1])

			left = last + 1
		}

	}

	return splitStrs
}

/*
题目描述
有N个文件，每个文件的编号从0至N-1，相应大小分别记为S(i)。给定磁盘空间为C，
试实现一个函数从N个文件中选出若干个连续的文件拷贝到磁盘中，使得磁盘剩余空间最小。
示例：C, S(i) 分别为7，[1, 2, 3, 5, 4]，可求得「起始编号」、「结束编号」为 0, 2。
*/
func FindMaxSeq(c int, sizes []int) []int {
	var left, right, maxNum int

	for i := 0; i < len(sizes); i++ {
		var sum int
		for j := 0; i+j < len(sizes); j++ {
			if j == 0 {
				sum = sizes[i]
			} else {
				sum = sum + sizes[i+j]
			}

			if sum > c {
				break
			}

			if sum > maxNum {
				maxNum = sum
				left, right = i, i+j
			}
		}
	}

	return []int{left, right}
}

// 字符的全排列问题: 如 "12" 排列有"1" "2" "12"
// TODO
func StrPermutations(s string) []string {
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

// assume no error
func StrCalculate(str string) float64 {

	arr := make([]string, 0)
	firstOps := make([]int, 0)

	var start int = 0
	for i := 0; i < len(str); i++ {
		if str[i] == byte('+') || str[i] == byte('-') || str[i] == byte('*') || str[i] == byte('/') {
			arr = append(arr, str[start:i])
			arr = append(arr, string(str[i]))
			start = i + 1

			if str[i] == byte('*') || str[i] == byte('/') {
				firstOps = append(firstOps, len(arr)-1)
			}

		}
	}
	arr = append(arr, str[start:])

	// * /
	index2val := make(map[int]float64)
	index2right := make(map[int]int)
	for _, index := range firstOps {
		var leftVal float64
		if val, exist := index2val[index-1]; exist {
			leftVal = val
		} else {
			a, _ := strconv.Atoi(arr[index-1])
			leftVal = float64(a)
		}

		b, _ := strconv.Atoi(arr[index+1])

		var result float64
		if arr[index] == "*" {
			result = leftVal * float64(b)
		} else {
			result = leftVal / float64(b)
		}
		index2right[index-1] = index + 1
		index2val[index+1] = result
	}

	var sum float64
	oldPos, leftPos := 0, 0
	for {
		if pos, exist := index2right[leftPos]; exist {
			leftPos = pos
			continue
		}
		break
	}
	if leftPos == oldPos {
		a, _ := strconv.Atoi(arr[oldPos])
		sum = float64(a)
	} else {
		sum = float64(index2val[leftPos])
	}

	for i := leftPos + 1; i+1 < len(arr); i++ {
		if arr[i] != "+" && arr[i] != "-" {
			continue
		}

		// right
		rightPos, oldPos := i+1, i+1
		for {
			if pos, exist := index2right[rightPos]; exist {
				rightPos = pos
				continue
			}
			break
		}
		var rightVal float64
		if rightPos == oldPos {
			b, _ := strconv.Atoi(arr[rightPos])
			rightVal = float64(b)
		} else {
			i = rightPos
			rightVal = index2val[rightPos]
		}

		if arr[i] == "+" {
			sum = sum + rightVal
		} else {
			sum = sum - rightVal
		}

	}

	return sum
}
