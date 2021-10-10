package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t, _ := strconv.Atoi(input.Text())
	for ; t > 0; t-- {
		input.Scan()
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}

