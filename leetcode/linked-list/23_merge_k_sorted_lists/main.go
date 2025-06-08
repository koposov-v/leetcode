package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}

	cur := dummyNode

	for {
		minIdx := getIndexMinValue(lists)
		if minIdx == nil {
			break
		}

		cur.Next = lists[*minIdx]
		cur = cur.Next
		lists[*minIdx] = lists[*minIdx].Next
	}

	return dummyNode.Next
}

func getIndexMinValue(lists []*ListNode) *int {
	minV := 1<<63 - 1
	idx := 0
	for i, list := range lists {
		if list == nil {
			continue
		}
		if list.Val < minV {
			minV = list.Val
			idx = i
		}
	}

	if minV == 1<<63-1 {
		return nil
	}

	return &idx
}
