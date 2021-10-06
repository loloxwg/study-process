package leetcode

import "math"

func maxSubArray(nums []int) int {
	dp:= make([]int,len(nums))
	dp[0]=nums[0]

	for i:=1 ;i<len(nums);i++{
		if dp[i-1]>0{
			dp[i]=dp[i-1]+nums[i]
		}else {
			dp[i]=nums[i]
		}
	}
	var res =dp[0]
	for i := 0; i <len(nums); i++ {
		res= int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}

//dp[i]：表示以 nums[i] 结尾 的 连续 子数组的最大和。