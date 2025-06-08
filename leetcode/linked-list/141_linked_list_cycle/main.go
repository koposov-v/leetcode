package main

import "fmt"

func main() {
	fmt.Println(4 + 2%10)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hashMap := make(map[*ListNode]struct{})

	cur := head

	for cur != nil {
		if _, ok := hashMap[cur]; !ok {
			hashMap[cur] = struct{}{}
		} else {
			return true
		}
		cur = cur.Next
	}

	return false
}
