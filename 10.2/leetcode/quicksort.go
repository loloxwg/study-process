package main

import (
	"sort"
)

func sortArray(nums []int) []int {
	sort.Sort(sort.IntSlice(nums))
	return nums
}

func sortArray2(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) []int {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		sink(nums, 0, end)
	}
	return nums
}

func sink(heap []int, root, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child < end && heap[child] <= heap[child+1] {
			child++
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j, base := left, right, nums[left]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func quick_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	nums[i], nums[(i+j)>>1] = nums[(i+j)>>1], nums[i]
	pivot := nums[i]
	for i < j {
		for i < j && pivot <= nums[j] {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quick_sort(nums, l, i-1)
	quick_sort(nums, i+1, r)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result
}
