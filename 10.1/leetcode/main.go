package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}
func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}


func a (s string)int{
	m:= map[byte]int{}
	fmt.Println("first result: ",m)
	n:=len(s)
	rk,ans:=0,0
	for i := 0; i < n; i++ {
		if i!=0{
			delete(m,s[i-1])
		}
		for rk<n &&m[s[rk]]==0{
			m[s[rk]]++//0变成1
			fmt.Println(m)
			rk++
		}
		ans=max(ans,rk-i)
	}
	return ans
}

func main() {
	s:="abcabcbb"
	b:=a(s)
	fmt.Println(b)

}