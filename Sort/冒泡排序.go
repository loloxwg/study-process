package main

import "fmt"

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
func bobSort(nums []int)[]int{
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] >nums[j]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	return nums
}
func bubbleSort2(arr []int) []int {
	swapped:=true
	for swapped {
		swapped=false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] >arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
				swapped=true
			}
		}
	}
	return arr
}

func main() {
	var nums =[]int{6,7,3,4,6,1, 2, 3, 4,4,4,4,4}
	//numsSort:=bobSort(nums)
	numsSort:=bubbleSort2(nums)
	fmt.Println("Bob sorting",numsSort)
}