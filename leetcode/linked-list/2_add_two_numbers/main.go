package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10
		sum %= 10

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		sum %= 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
