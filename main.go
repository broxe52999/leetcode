package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

// 轮转数组
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// H指数
func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if citations[i] >= n-i {
			ans = n - i
		}
	}
	return ans
}
