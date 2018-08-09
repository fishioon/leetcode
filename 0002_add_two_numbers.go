package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l := new(ListNode)
	node := l
	emptyNode := new(ListNode)
	more := 0
	for l1 != nil || l2 != nil || more > 0 {
		if l1 == nil {
			l1 = emptyNode
		}
		if l2 == nil {
			l2 = emptyNode
		}
		num := l1.Val + l2.Val + more
		newNode := &ListNode{
			Val: num % 10,
		}
		more = num / 10
		node.Next = newNode
		node = newNode
		l1 = l1.Next
		l2 = l2.Next
	}
	return l.Next
}
