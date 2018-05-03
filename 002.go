package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	l4 := l3
	a := 0
	for {
		l4.Val = a
		if l1 != nil {
			l4.Val += l1.Val
		}
		if l2 != nil {
			l4.Val += l2.Val
		}
		if l4.Val > 9 {
			a = l4.Val / 10
			l4.Val = l4.Val % 10
		} else {
			a = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			if a > 0 {
				l4.Next = new(ListNode)
				l4.Next.Val = a
			}
			break
		} else {
			l4.Next = new(ListNode)
			l4 = l4.Next
		}
	}
	return l3
}
