/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : tree_test.go
*   coder: zemanzeng
*   date : 2022-02-02 10:26:36
*   desc :
*
================================================================*/

package tree

import (
	"reflect"
	"testing"
)

func listInterfaceEqual(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if reflect.TypeOf(a[i]) != reflect.TypeOf(b[i]) ||
			reflect.ValueOf(a[i]) != reflect.ValueOf(b[i]) {
			return false
		}
	}
	return true
}

type orderResult struct {
	input  []interface{}
	output []interface{}
}

func TestPreorderTraversal(t *testing.T) {

	results := []*orderResult{
		{
			input:  []interface{}{1, nil, 2, nil, nil, 3},
			output: []interface{}{1, 2, 3},
		},
		{

			input:  []interface{}{1},
			output: []interface{}{1},
		},
		{

			input:  []interface{}{},
			output: []interface{}{},
		},
	}

	for _, result := range results {
		root := NewTreeNodes(result.input...)
		orderList := PreorderTraversal(root)
		if !listInterfaceEqual(orderList, result.output) {
			t.Errorf("pre_order_traversal not qual. input:%+v orderlist:%+v output:%v",
				result.input, orderList, result.output)
		}
	}

}

func TestInorderTraversal(t *testing.T) {
	results := []*orderResult{
		{
			input:  []interface{}{1, nil, 2, nil, nil, 3},
			output: []interface{}{1, 3, 2},
		},
		{

			input:  []interface{}{1},
			output: []interface{}{1},
		},
		{

			input:  []interface{}{},
			output: []interface{}{},
		},
	}

	for _, result := range results {
		root := NewTreeNodes(result.input...)
		orderList := InorderTraversal(root)
		if !listInterfaceEqual(orderList, result.output) {
			t.Errorf("in_order_traversal not qual. input:%+v orderlist:%+v output:%v",
				result.input, orderList, result.output)
		}
	}

}

func TestPostorderTraversal(t *testing.T) {

	results := []*orderResult{
		{
			input:  []interface{}{1, nil, 2, nil, nil, 3},
			output: []interface{}{3, 2, 1},
		},
		{

			input:  []interface{}{1},
			output: []interface{}{1},
		},
		{

			input:  []interface{}{},
			output: []interface{}{},
		},
	}

	for _, result := range results {
		root := NewTreeNodes(result.input...)
		orderList := PostorderTraversal(root)
		if !listInterfaceEqual(orderList, result.output) {
			t.Errorf("post_order_traversal not qual. input:%+v orderlist:%+v output:%v",
				result.input, orderList, result.output)
		}
	}

}
