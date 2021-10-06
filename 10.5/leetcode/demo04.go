package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
//排序
func merge2(nums1 []int, m int, nums2 []int, n int) {
	var merged=make([]int,m+n)
	merged=nums1[:m]
	//for i := 0; i < n; i++ {
	//	merged=append(merged,nums2[i])
	//}
	merged=append(merged,nums2[:n]...)
	sort.Sort(sort.IntSlice(merged))
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for{
		if p1==m{
			sorted=append(sorted,nums2[p2:]...)
			break
		}
		if p2==n{
			sorted=append(sorted,nums1[p1:]...)
			break
		}
		if nums1[p1]<nums2[p2]{
			sorted=append(sorted,nums1[p1])
			p1++
		}else {
			sorted=append(sorted,nums2[p2])
			p2++
		}
	}
	copy(nums1,sorted)
}