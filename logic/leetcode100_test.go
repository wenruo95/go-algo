package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	type testData struct {
		p    []interface{}
		q    []interface{}
		same bool
	}

	datas := []*testData{
		{
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 3},
			same: true,
		},
		{
			p:    []interface{}{1, 2},
			q:    []interface{}{1, nil, 2},
			same: false,
		},

		{
			p:    []interface{}{1, 2, 1},
			q:    []interface{}{1, 1, 2},
			same: false,
		},
	}

	for _, data := range datas {
		p := genTreeNode(data.p)
		q := genTreeNode(data.q)
		same := IsSameTree(p, q)
		assert.Equal(t, data.same, same)
	}
}
