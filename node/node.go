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

func NewCircle(pos int, a ...interface{}) *ListNode {
	if pos < 0 {
		return New(a...)
	}

	list := make([]*ListNode, len(a))
	for i := 0; i < len(a); i++ {
		list[i] = &ListNode{Val: a[i]}

		if i > 0 {
			list[i-1].Next = list[i]
		}
	}

	list[len(list)-1].Next = list[pos]
	return list[0]
}

func (n *ListNode) Dump() []interface{} {
	values := make([]interface{}, 0)

	node := n
	for i := 0; i < 1000 && node != nil; i++ {
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

// 链表反转
// a -> b -> c -> d
// a -> ( b -> c -> d )
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

// 反转前N个节点
// a -> b -> c -> d -> e
// a -> (b -> c -> d) -> e
// (b -> c -> d) -> a -> suffix
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

// 反转[left,right]中节点
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

// 优化版本
func ListNodeReorder2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	list := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		list = append(list, node)
	}

	low, high := 0, len(list)-1
	for low < high {
		list[low].Next = list[high]
		low = low + 1
		if low == high {
			break
		}

		list[high].Next = list[low]
		high = high - 1
	}
	list[low].Next = nil
	return list[0]
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

// leetcode 24: https://leetcode.com/problems/swap-nodes-in-pairs/
func SwapPairs(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode = head
	for next != nil && next.Next != nil {
		first, second := next, next.Next
		first.Next = second.Next
		second.Next = first

		if prev == nil {
			head = second
		} else {
			prev.Next = second
		}
		prev = first
		next = prev.Next
	}

	return head
}

// leetcode 25: https://leetcode.com/problems/reverse-nodes-in-k-group/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	list := make([]*ListNode, 0)
	var node *ListNode = head
	for node != nil {
		list = append(list, node)
		node = node.Next
	}

	cnt := len(list) / k
	for i := 0; i < cnt; i++ {
		low, high := i*k, (i+1)*k-1
		for j := 0; j < k/2; j++ {
			list[low+j], list[high-j] = list[high-j], list[low+j]
		}
	}

	head = list[0]
	for i := 0; i < cnt*k && i+1 < len(list); i++ {
		list[i].Next = list[i+1]
	}

	return head
}

func ReverseKGroup2(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	var last *ListNode
	var node *ListNode = head
	for {
		list := make([]*ListNode, 0)
		for i := 0; i < k && node != nil; i++ {
			list = append(list, node)
			node = node.Next
		}
		if len(list) < k {
			break
		}

		low, high := 0, k-1
		for i := 0; i < k/2; i++ {
			list[low+i], list[high-i] = list[high-i], list[low+i]
		}
		for i := 0; i < high; i++ {
			list[i].Next = list[i+1]
		}
		list[high].Next = node

		if last == nil {
			head = list[low]
		} else {
			last.Next = list[low]
		}

		last = list[high]
	}

	return head
}

func ReverseKGroup3(head *ListNode, k int) *ListNode {

	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	var node *ListNode = head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead := reverse(head, node)
	head.Next = ReverseKGroup3(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	var node *ListNode = last
	for first != last {
		tmp := first.Next

		first.Next = node
		node = first

		first = tmp
	}
	return node
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		if fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

// leetcode 2: https://leetcode.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nextNum int

	var head, prev *ListNode
	for l1 != nil || l2 != nil {
		var num int = nextNum
		if l1 != nil {
			num = num + l1.Val.(int)
		}
		if l2 != nil {
			num = num + l2.Val.(int)
		}
		nextNum = num / 10

		node := &ListNode{Val: num % 10}
		if prev == nil {
			prev = node
			head = node
		} else {
			prev.Next = node
			prev = node
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if nextNum > 0 {
		node := &ListNode{Val: nextNum}
		prev.Next = node
	}

	return head
}

// leetcode 82: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
func DeleteDuplicates(head *ListNode) *ListNode {
	var prev, rhead, node *ListNode
	for head != nil {
		if (prev != nil && prev.Val.(int) == head.Val.(int)) ||
			(head.Next != nil && head.Next.Val.(int) == head.Val.(int)) {

			prev = head
			head = head.Next
			continue
		}
		prev = head

		if rhead == nil {
			rhead = head
			node = head
			continue
		}
		node.Next = head
		node = node.Next
		head = head.Next

	}
	if node != nil {
		node.Next = nil
	}

	return rhead
}

// leetcode 61: https://leetcode.com/problems/rotate-list/
func RotateRight(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil || head.Next == nil {
		return head
	}

	// find n-th
	var fast *ListNode
	for i := 1; i <= k; i++ {
		if i == 1 {
			fast = head
			continue
		}

		fast = fast.Next

		if fast.Next == nil {
			if k%i == 0 {
				return head
			}
			k = i + k%i

			i = i + 1
			fast = head
		}

	}

	var slow *ListNode
	for fast.Next != nil {
		if slow == nil {
			slow = head
			fast = fast.Next
			continue
		}

		slow = slow.Next
		fast = fast.Next
	}

	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}

// leetcode 83: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func DeleteDuplicates2(head *ListNode) *ListNode {
	var current *ListNode = head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// leetcode 86: https://leetcode.com/problems/partition-list/
func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var smallerHead, smallerNode, largerHead *ListNode
	if head.Val.(int) >= x {
		largerHead = head
	} else {
		smallerHead = head
		smallerNode = head
	}

	for head != nil && head.Next != nil {
		node := head.Next
		if node.Val.(int) >= x {
			if largerHead == nil {
				largerHead = node
			}
			head = node
			continue
		}

		if smallerHead == nil {
			smallerHead = node
			smallerNode = node
		}

		if largerHead == nil {
			smallerNode = node
			head = node
			continue
		}
		head.Next = node.Next
		node.Next = nil
		smallerNode.Next = node
		smallerNode = node
	}

	if smallerHead == nil {
		return largerHead
	}
	smallerNode.Next = largerHead
	return smallerHead
}

func Partition2(head *ListNode, x int) *ListNode {
	var node, smallerNode *ListNode
	if node = head; node != nil && node.Val.(int) < x {
		smallerNode = node
	}

	for node != nil && node.Next != nil {
		if node.Next.Val.(int) >= x {
			node = node.Next
			continue
		}

		if smallerNode == nil {
			smallerNode = node.Next
			node.Next = node.Next.Next
			smallerNode.Next = head
			head = smallerNode
			continue
		}
		if smallerNode == node {
			node = node.Next
			smallerNode = node
			continue
		}

		tmp := node.Next
		node.Next = node.Next.Next
		tmp.Next = smallerNode.Next
		smallerNode.Next = tmp
		smallerNode = tmp
	}

	return head
}

func Partition3(head *ListNode, x int) *ListNode {
	var (
		smallHead = new(ListNode)
		largeHead = new(ListNode)
		small     = smallHead
		large     = largeHead
		node      = head
	)
	for node != nil {
		if node.Val.(int) < x {
			small.Next = node
			node = node.Next

			small = small.Next
			small.Next = nil
		} else {
			large.Next = node
			node = node.Next

			large = large.Next
			large.Next = nil
		}
	}
	small.Next = largeHead.Next
	return smallHead.Next
}

// leetcode 92: https://leetcode.com/problems/reverse-linked-list-ii/
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	var suffix *ListNode
	var reversefn func(node *ListNode, n int) *ListNode
	reversefn = func(node *ListNode, n int) *ListNode {
		if node == nil || node.Next == nil {
			suffix = nil
			return node
		}
		if n <= 1 {
			suffix = node.Next
			return node
		}

		last := reversefn(node.Next, n-1)
		node.Next.Next = node
		node.Next = suffix
		return last
	}

	if left == 1 {
		return reversefn(head, right)
	}
	head.Next = ReverseBetween(head.Next, left-1, right-1)

	return head
}
