package logic

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func randIntList(size int, max int) []int {
	arr := make([]int, 0)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(max))
	}
	return arr
}

func Test_QuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	maxList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 100000}
	sizeList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200}
	for _, size := range sizeList {
		for _, max := range maxList {
			result := randIntList(size, max)

			result2 := make([]int, size)
			copy(result2, result)

			expect := make([]int, size)
			copy(expect, result)
			sort.Ints(expect)

			QuickSortRecur(result)
			QuickSortItera(result2)

			assert.Equal(t, expect, result)
			assert.Equal(t, expect, result2)
		}
	}

}
