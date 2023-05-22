package logic

import (
	"container/heap"
)

type IntListHeap [][]int

func (h IntListHeap) Len() int            { return len(h) }
func (h IntListHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h IntListHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntListHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func TrapRainWater(heightMap [][]int) int {
	M, N := len(heightMap), len(heightMap[0])
	visited := make([][]bool, M)
	for i := 0; i < M; i++ {
		visited[i] = make([]bool, N)
	}

	h := make(IntListHeap, 0)
	for row := 0; row < M; row++ {
		heap.Push(&h, []int{heightMap[row][0], row, 0})
		heap.Push(&h, []int{heightMap[row][N-1], row, N - 1})
		visited[row][0] = true
		visited[row][N-1] = true
	}
	for col := 0; col < N; col++ {
		heap.Push(&h, []int{heightMap[0][col], 0, col})
		heap.Push(&h, []int{heightMap[M-1][col], M - 1, col})
		visited[0][col] = true
		visited[M-1][col] = true
	}

	var vol int

	// up down left right
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for h.Len() > 0 {
		item := heap.Pop(&h).([]int)
		height, row, col := item[0], item[1], item[2]

		for _, dir := range directions {
			r, c := row+dir[0], col+dir[1]
			if r >= 0 && r < M && c >= 0 && c < N && !visited[r][c] {
				visited[r][c] = true
				if heightMap[r][c] < height {
					vol = vol + height - heightMap[r][c]
				}
				heap.Push(&h, []int{maxInt(height, heightMap[r][c]), r, c})
			}
		}
	}

	return vol
}
