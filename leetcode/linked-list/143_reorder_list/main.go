package main

import "fmt"

func main() {
	fmt.Println("lOL")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := preMid(head)
	p2 := reverse(mid.Next)

	mid.Next = nil

	p1 := head

	for p2 != nil {
		p1Next := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1Next
	}
}

func preMid(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
