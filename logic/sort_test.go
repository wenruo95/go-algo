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

func testSort(t *testing.T, name string, sortFn func(arr []int), debug ...bool) {
	rand.Seed(time.Now().UnixNano())

	maxList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 100000}
	sizeList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 200}

	if len(debug) > 0 && debug[0] {
		maxList = []int{8}
		sizeList = []int{10}
	}
	for _, size := range sizeList {
		for _, max := range maxList {
			result := randIntList(size, max)

			expect := make([]int, size)
			copy(expect, result)
			sort.Ints(expect)

			sortFn(result)

			if !assert.Equal(t, expect, result) {
				t.Errorf("assert %v error", name)
			}
		}
	}

}

func Test_QuickSort(t *testing.T) {
	testSort(t, "quick_sort_recursion", QuickSortRecur)
	testSort(t, "quick_sort_iteration", QuickSortItera)
}

func Test_MergeSort(t *testing.T) {
	testSort(t, "merge_sort_recursion", MergeSortRecur)
	testSort(t, "merge_sort_iteration", MergeSortItera)
}
