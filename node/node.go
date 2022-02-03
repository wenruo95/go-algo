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

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func New(a ...interface{}) *ListNode {
	if len(a) == 0 {
		return nil
	}

	head := new(ListNode)

	node := head
	for index, v := range a {
		if index != 0 {
			node.Next = new(ListNode)
			node = node.Next
		}
		node.Val = v
	}
	return head
}

func (n *ListNode) Dump() []interface{} {
	values := make([]interface{}, 0)

	node := n
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return values
}

func (n *ListNode) Equal(src *ListNode) bool {
	dst := n
	for dst != nil && src != nil {
		if dst.Val != src.Val {
			return false
		}
		dst = dst.Next
		src = src.Next
	}
	return (dst == nil) && (src == nil)
}

func ListNodeReverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := ListNodeReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var suffix *ListNode

func ListNodeReverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		suffix = nil
		return head
	}
	if n <= 1 {
		suffix = head.Next
		return head
	}
	last := ListNodeReverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = suffix
	return last
}

func ListNodeReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return ListNodeReverseN(head, right)
	}

	head.Next = ListNodeReverseBetween(head.Next, left-1, right-1)
	return head
}

// leetcode 143: https://leetcode.com/problems/reorder-list/
func ListNodeReorder(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		node.Next = ListNodeReverse(node.Next)
		node = node.Next
	}
	return head
}

// leetcode 19: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 1; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head
	}

	if fast.Next == nil {
		if fast == head {
			return nil
		}

		deleted := head
		head = head.Next
		deleted.Next = nil

		return head
	}

	var slow *ListNode
	for fast.Next != nil {
		fast = fast.Next
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
	}

	deleted := slow.Next
	slow.Next = deleted.Next
	deleted.Next = nil
	return head
}

// leetcode 21: https://leetcode.com/problems/merge-two-sorted-lists/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head, node *ListNode
	for list1 != nil && list2 != nil {
		var curNode *ListNode
		if list1.Val.(int) < list2.Val.(int) {
			curNode = list1
			list1 = list1.Next
		} else {
			curNode = list2
			list2 = list2.Next
		}

		if head == nil {
			head = curNode
			node = curNode
			continue
		}

		node.Next = curNode
		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
		node = node.Next
	}
	if list2 != nil {
		node.Next = list2
		node = node.Next
	}

	return head
}

// leetcode 23: https://leetcode.com/problems/merge-k-sorted-lists/
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	node := lists[0]
	for i := 1; i < len(lists); i++ {
		node = MergeTwoLists(node, lists[i])
	}
	return node
}
