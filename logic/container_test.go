package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainerStack(t *testing.T) {
	st := NewStack()

	st.Push(100)
	assert.Equal(t, 1, st.Size())

	st.Push(200)
	assert.Equal(t, 2, st.Size())

	v, ok := st.Top().(int)
	assert.True(t, ok)
	assert.Equal(t, 200, v)
	assert.Equal(t, 2, st.Size())

	v2, ok := st.Pop().(int)
	assert.True(t, ok)
	assert.Equal(t, 200, v2)
	assert.Equal(t, 1, st.Size())
	assert.Equal(t, 100, st.Top().(int))

	st.Push(300)
	assert.Equal(t, 2, st.Size())
	assert.Equal(t, 300, st.Top().(int))
	st.Push(400)
	assert.Equal(t, 3, st.Size())
	assert.Equal(t, 400, st.Top().(int))
	st.Push(500)
	assert.Equal(t, 4, st.Size())
	assert.Equal(t, 500, st.Top().(int))

	assert.Equal(t, 500, st.Pop().(int))
	assert.Equal(t, 400, st.Pop().(int))
	assert.Equal(t, 300, st.Pop().(int))
	assert.Equal(t, 100, st.Pop().(int))

	assert.Equal(t, st.Top(), nil)
	assert.Equal(t, st.Size(), 0)
}

func Test_ContainerQueue(t *testing.T) {
	queue := NewQueue()

	queue.Push(100)
	assert.Equal(t, 1, queue.Size())

	queue.Push(200)
	assert.Equal(t, 2, queue.Size())

	v, ok := queue.Top().(int)
	assert.True(t, ok)
	assert.Equal(t, 100, v)
	assert.Equal(t, 2, queue.Size())

	v2, ok := queue.Pop().(int)
	assert.True(t, ok)
	assert.Equal(t, 100, v2)
	assert.Equal(t, 1, queue.Size())
	assert.Equal(t, 200, queue.Top().(int))

	queue.Push(300)
	assert.Equal(t, 2, queue.Size())
	assert.Equal(t, 200, queue.Top().(int))
	queue.Push(400)
	assert.Equal(t, 3, queue.Size())
	assert.Equal(t, 200, queue.Top().(int))
	queue.Push(500)
	assert.Equal(t, 4, queue.Size())
	assert.Equal(t, 200, queue.Top().(int))

	assert.Equal(t, 200, queue.Pop().(int))
	assert.Equal(t, 300, queue.Pop().(int))
	assert.Equal(t, 400, queue.Pop().(int))
	assert.Equal(t, 500, queue.Pop().(int))

	assert.Equal(t, queue.Top(), nil)
	assert.Equal(t, queue.Size(), 0)
}
