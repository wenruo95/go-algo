/*================================================================
*   Copyright (C) 2022. All rights reserved.
*
*   file : node.go
*   coder: zemanzeng
*   date : 2022-01-20 17:05:06
*   desc : 单链表
*
================================================================*/

package reverse

type Node struct {
	Value interface{}
	Next  *Node
}

func New(a ...interface{}) *Node {
	if len(a) == 0 {
		return nil
	}

	head := new(Node)

	node := head
	for index, v := range a {
		if index != 0 {
			node.Next = new(Node)
			node = node.Next
		}
		node.Value = v
	}
	return head
}

func (n *Node) Dump() []interface{} {
	values := make([]interface{}, 0)

	node := n
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

func (n *Node) Equal(src *Node) bool {
	dst := n
	for dst != nil && src != nil {
		if dst.Value != src.Value {
			return false
		}
		dst = dst.Next
		src = src.Next
	}
	return (dst == nil) && (src == nil)
}

func NodeReverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	last := NodeReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var suffix *Node

func NodeReverseN(head *Node, n int) *Node {
	if head == nil || head.Next == nil {
		suffix = nil
		return head
	}
	if n <= 1 {
		suffix = head.Next
		return head
	}
	last := NodeReverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = suffix
	return last
}

func NodeReverseBetween(head *Node, left int, right int) *Node {
	if left == 1 {
		return NodeReverseN(head, right)
	}

	head.Next = NodeReverseBetween(head.Next, left-1, right-1)
	return head
}
