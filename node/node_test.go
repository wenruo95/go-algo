/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : node_test.go
*   coder: zemanzeng
*   date : 2022-01-20 17:11:11
*   desc :
*
================================================================*/

package reverse

import (
	"testing"
)

func TestEquals(t *testing.T) {
	type equalResult struct {
		a      []interface{}
		b      []interface{}
		result bool
	}

	results := []*equalResult{
		{
			a:      []interface{}{1, 2},
			b:      []interface{}{1},
			result: false,
		},
		{
			a:      []interface{}{1, 2},
			b:      []interface{}{1, 2},
			result: true,
		},
		{
			a:      []interface{}{1, 2, 3},
			b:      []interface{}{1},
			result: false,
		},
		{
			a:      []interface{}{1, 2, 3, 4},
			b:      []interface{}{1, 2, 3, 4},
			result: true,
		},
		{
			a:      []interface{}{1},
			b:      []interface{}{1},
			result: true,
		},
		{
			a:      []interface{}{},
			b:      []interface{}{},
			result: true,
		},
	}

	for _, result := range results {
		n1 := New(result.a...)
		n2 := New(result.b...)
		if equal := n1.Equal(n2); equal != result.result {
			t.Errorf("node not equal. a:%+v b:%+v equal:%v result:%v",
				result.a, result.b, equal, result.result)
		}
	}

}

func listEqual(dst, src []interface{}) bool {
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

func TestDump(t *testing.T) {
	type dumpResult struct {
		a []interface{}
	}
	results := []*dumpResult{
		{
			a: []interface{}{},
		},

		{
			a: []interface{}{1},
		},

		{
			a: []interface{}{1, 2},
		},

		{
			a: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, result := range results {
		dump := New(result.a...).Dump()
		if !listEqual(dump, result.a) {
			t.Errorf("dump:%+v not equal result:%+v", dump, result)
		}
	}

}

func listReverse(l []interface{}) []interface{} {
	r := make([]interface{}, 0)
	for i := len(l) - 1; i >= 0; i-- {
		r = append(r, l[i])
	}
	return r
}

func TestReverse(t *testing.T) {
	type reverseResult struct {
		a []interface{}
	}
	results := []*reverseResult{
		{
			a: []interface{}{},
		},

		{
			a: []interface{}{1},
		},

		{
			a: []interface{}{1, 2},
		},

		{
			a: []interface{}{1, 2, 3, 4, 5},
		},
	}

	for _, result := range results {
		n1 := New(result.a...)
		l1 := n1.Dump()

		n2 := ListNodeReverse(n1)
		l2 := n2.Dump()
		l3 := listReverse(l2)

		if !listEqual(l1, l3) {
			t.Errorf("reverse error. l1:%+v l2:%+v l3:%+v", l1, l2, l3)
		}
	}
}

func TestReverseN(t *testing.T) {
	type reverseNResult struct {
		a []interface{}
		n int
		b []interface{}
	}
	results := []*reverseNResult{
		{
			a: []interface{}{},
			n: 0,
			b: []interface{}{},
		},
		{
			a: []interface{}{1, 2},
			n: 1,
			b: []interface{}{1, 2},
		},
		{
			a: []interface{}{1, 2, 3, 4, 5},
			n: 2,
			b: []interface{}{2, 1, 3, 4, 5},
		},
		{
			a: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			n: 6,
			b: []interface{}{6, 5, 4, 3, 2, 1, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			a: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			n: 15,
			b: []interface{}{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, result := range results {
		n1 := New(result.a...)
		l1 := n1.Dump()

		n2 := ListNodeReverseN(n1, result.n)
		l2 := n2.Dump()

		if !listEqual(l2, result.b) {
			t.Errorf("reverse error. l1:%+v n:%v l2:%+v l3:%+v", l1, result.n, l2, result.b)
		}

	}

}

func TestReverseBetween(t *testing.T) {
	type reverseNResult struct {
		a     []interface{}
		left  int
		right int
		b     []interface{}
	}
	results := []*reverseNResult{
		{
			a:     []interface{}{1, 2},
			left:  1,
			right: 2,
			b:     []interface{}{2, 1},
		},
		{
			a:     []interface{}{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
			b:     []interface{}{1, 4, 3, 2, 5},
		},
		{
			a:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			left:  2,
			right: 10,
			b:     []interface{}{1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 11, 12, 13, 14, 15},
		},
		{
			a:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			left:  12,
			right: 15,
			b:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 15, 14, 13, 12},
		},
	}

	for _, result := range results {
		n1 := New(result.a...)
		l1 := n1.Dump()

		n2 := ListNodeReverseBetween(n1, result.left, result.right)
		l2 := n2.Dump()

		if !listEqual(l2, result.b) {
			t.Errorf("reverse error. l1:%+v left:%v right:%v l2:%+v l3:%+v",
				l1, result.left, result.right, l2, result.b)
		}

	}

}

func TestListNodeReorder(t *testing.T) {
	type reorderResult struct {
		a []interface{}
		b []interface{}
	}

	results := []*reorderResult{
		{
			a: []interface{}{1, 2, 3, 4},
			b: []interface{}{1, 4, 2, 3},
		},
		{
			a: []interface{}{1, 2, 3, 4, 5},
			b: []interface{}{1, 5, 2, 4, 3},
		},
		{
			a: []interface{}{1},
			b: []interface{}{1},
		},
		{
			a: []interface{}{1, 2},
			b: []interface{}{1, 2},
		},
		{
			a: []interface{}{1, 2, 3},
			b: []interface{}{1, 3, 2},
		},
		{
			a: []interface{}{},
			b: []interface{}{},
		},
	}

	for _, result := range results {
		reorder := ListNodeReorder(New(result.a...))
		if equal := reorder.Equal(New(result.b...)); !equal {
			t.Errorf("node not equal. a:%+v b:%+v reorder:%+v equal:%v",
				result.a, result.b, reorder, equal)
		}
	}

}

func TestRemoveNThFromEnd(t *testing.T) {
	type removeResult struct {
		a []interface{}
		n int
		b []interface{}
	}

	results := []*removeResult{
		{
			a: []interface{}{1, 2, 3, 4, 5},
			n: 2,
			b: []interface{}{1, 2, 3, 5},
		},
		{
			a: []interface{}{1},
			n: 1,
			b: []interface{}{},
		},
		{
			a: []interface{}{1, 2},
			n: 1,
			b: []interface{}{1},
		},
		{
			a: []interface{}{1, 2},
			n: 2,
			b: []interface{}{2},
		},
		{
			a: []interface{}{1, 2, 3},
			n: 3,
			b: []interface{}{2, 3},
		},
	}

	for _, result := range results {
		removedNode := RemoveNthFromEnd(New(result.a...), result.n)
		if equal := removedNode.Equal(New(result.b...)); !equal {
			t.Errorf("remove_n_th_from_end. result:%+v removed:%+v equal:%v",
				result, removedNode.Dump(), equal)
		}
	}

}
