package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
    cur := dummy

    for list1 != nil || list2 != nil {
        if getVal(list1) < getVal(list2) {
            cur.Next = list1
            list1 = list1.Next
        } else {
            cur.Next = list2
            list2 = list2.Next
        }

        cur = cur.Next
    }

	return dummy.Next
}

func getVal(head *ListNode) int {
    if head == nil {
        return 1 << 63 - 1
    }
    
    return head.Val
}

func main() {
	//ans := mergeTwoLists(list1, list2)
}
