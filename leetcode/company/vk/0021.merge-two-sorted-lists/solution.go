// Created by Vyacheslav Koposov at 2025/06/08 13:06
// leetgo: 1.4.13
// https://leetcode.cn/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	list1 := Deserialize[*ListNode](ReadLine(stdin))
	list2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}
