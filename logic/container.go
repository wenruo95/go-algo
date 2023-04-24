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
