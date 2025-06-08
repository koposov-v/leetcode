// Created by Vyacheslav Koposov at 2025/06/08 13:06
// leetgo: 1.4.13
// https://leetcode.cn/problems/search-in-rotated-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
