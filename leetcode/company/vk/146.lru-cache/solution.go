// Created by Vyacheslav Koposov at 2025/06/08 13:07
// leetgo: 1.4.13
// https://leetcode.cn/problems/lru-cache/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type LRUCache struct {
    
}

func Constructor(capacity int) LRUCache {

	return LRUCache{}
}

func (l *LRUCache) Get(key int) (ans int) {

	return
}

func (l *LRUCache) Put(key int, value int)  {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	constructorParams := MustSplitArray(params[0])
	capacity := Deserialize[int](constructorParams[0])
	obj := Constructor(capacity)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "get":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Get(key))
			output = append(output, ans)
		case "put":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			value := Deserialize[int](methodParams[1])
			obj.Put(key, value)
			output = append(output, "null")
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
