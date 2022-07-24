package logic

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (st *Stack) Push(i int) {
	st.l.PushBack(i)
}

func (st *Stack) Pop() int {
	node := st.l.Back()
	st.l.Remove(node)
	return node.Value.(int)
}

func (st *Stack) Top() int {
	return st.l.Back().Value.(int)
}

func (st *Stack) Empty() bool {
	return st.l.Len() == 0
}
