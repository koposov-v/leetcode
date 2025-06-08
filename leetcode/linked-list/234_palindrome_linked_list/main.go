package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func isPalindrome(head *ListNode) bool {
	rightPart := halfNode(head)
	reverseRightPart := reverse(rightPart)

	first, right := head, reverseRightPart

	for first != nil && right != nil {
		if first.Val != right.Val {
			return false
		}

		first = first.Next
		right = right.Next
	}

	return true
}

func halfNode(head *ListNode) *ListNode {
	fast, slow := head, head
	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	var tail *ListNode

	cur := head

	for cur != nil {
		newNode := &ListNode{Val: cur.Val}
		if result == nil {
			result = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		cur = cur.Next
	}

	cur = result

	var prev *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
