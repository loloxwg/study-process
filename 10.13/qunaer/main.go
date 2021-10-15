package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)
func frequencySort2(s string) string {
	// 将字符串s转换为切片t1
	t1 := []byte(s)
	// 初始化映射表 table
	table := make(map[byte]int)
	for _, v := range t1 {
		table[v]++
	}
	// 进行排序
	t2 := make([]byte, 0)
	for k, v := range table {
		// 寻找插入的位置：index
		var index int
		for i := 0; i <= len(t2); i += table[t2[i]] {
			if i == len(t2) || table[t2[i]] <= v {
				index = i
				break
			}
		}
		// 执行插入
		for i := 0; i < v; i++ {
			t2 = append(t2[:index], append([]byte{k}, t2[index:]...)...)
		}
	}
	return string(t2)
}


func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}

	type pair struct {
		ch  byte
		cnt int
	}
	pairs := make([]pair, 0, len(cnt))
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}

func main() {
	a:=""
	b:=""
	for {
		_, err := fmt.Scanf("%s\n%s\n",&a,&b)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			c:=reoder(a,b)
			fmt.Printf("%s\n",c )
		}
	}
}

