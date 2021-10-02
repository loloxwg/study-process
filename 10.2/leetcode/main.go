package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	return nums[k-1]
}

func main() {
	nums := []int{3, 3,3,3,2, 1, 5, 6, 4}
	a := findKthLargest(nums,5)
	fmt.Println(a)
}
