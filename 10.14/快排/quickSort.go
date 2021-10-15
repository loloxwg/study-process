package main

func sortArray(nums []int) []int {
	quickSort(nums,0,len(nums)-1)
	return nums
}

func quickSort(nums []int ,left,right int){
	if left>=right{
		return
	}
	i,j:=left,right
	nums[i], nums[(i+j)>>1] = nums[(i+j)>>1], nums[i]
	p:=nums[i]
	for i<j{
		for i<j&&nums[j]>=p{
			j--
		}
		nums[i]=nums[j]
		for i<j&& nums[i]<=p{
			i++
		}
		nums[j]=nums[i]
	}
	nums[i]=p
	quickSort(nums,left,i-1)
	quickSort(nums,i+1,right)
}

func sortArray2(nums []int)[]int{
	quickSort(nums,0,len(nums)-1)
	return nums
}
func quickSort2(nums []int,left int,right int){
	if left>=right{
		return
	}
	i,j:=left,right
	nums[i],nums[(i+j)>>1]=nums[(i+j)>>1],nums[i]
	p:=nums[i]
	for i<j{
		for i<j &&nums[j]>=p{
			j--
		}
		nums[i]=nums[j]
		for j<j &&nums[j]<=p{
			i++
		}
		nums[j]=nums[i]
	}
	nums[i]=p
	quickSort(nums,left,i-1)
	quickSort(nums,i+1,right)
}

func quickSort3(nums []int,left int,right int){
	if left>=right{
		return
	}
	i,j:=left,right
	nums[i],nums[(i+j)>>1]=nums[(i+j)>>1],nums[i]
	p:=nums[i]
	for i<j {
		for i<j &&nums[j]>=p{
			j--
		}
		nums[i]=nums[j]
		for i<j &&nums[j]<=p{
			i++
		}
		nums[j]=nums[i]
	}
	nums[i]=p
	quickSort(nums,left,i-1)
	quickSort(nums,i+1,right)
}