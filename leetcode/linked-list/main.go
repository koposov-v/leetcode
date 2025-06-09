package ll

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateLinked(list []int) *ListNode {
	res := &ListNode{}
	for _, v := range list {
		res.Next = &ListNode{
			Val: v,
		}
		res = res.Next
	}

	return res.Next
}
