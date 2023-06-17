package logic

import "container/list"

type Stack interface {
	Top() interface{}
	Push(item interface{})
	Pop() interface{}
	Size() int
}

type stack struct {
	l *list.List
}

func NewStack() Stack {
	return &stack{l: list.New()}
}

func (st *stack) Top() interface{} {
	node := st.l.Back()
	if node == nil {
		return nil
	}
	return node.Value
}

func (st *stack) Push(item interface{}) {
	st.l.PushBack(item)
}

func (st *stack) Pop() interface{} {
	node := st.l.Back()
	if node == nil {
		return nil
	}
	st.l.Remove(node)
	return node.Value
}

func (st *stack) Size() int {
	return st.l.Len()
}

type queue struct {
	l *list.List
}

type Queue interface {
	Top() interface{}
	Push(item interface{})
	Pop() interface{}
	Size() int
}

func NewQueue() Queue {
	return &queue{l: list.New()}
}

func (qu *queue) Top() interface{} {
	node := qu.l.Front()
	if node == nil {
		return nil
	}
	return node.Value
}

func (qu *queue) Push(item interface{}) {
	qu.l.PushBack(item)
}

func (qu *queue) Pop() interface{} {
	node := qu.l.Front()
	if node == nil {
		return nil
	}
	qu.l.Remove(node)
	return node.Value
}

func (qu *queue) Size() int {
	return qu.l.Len()
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FrequentData struct {
	Count int
	Value int
}
type FrequentHeap []*FrequentData

func (h FrequentHeap) Len() int           { return len(h) }
func (h FrequentHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h FrequentHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FrequentHeap) Push(x interface{}) {
	*h = append(*h, x.(*FrequentData))
}

func (h *FrequentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
