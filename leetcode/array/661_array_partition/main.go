package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{4, 2, 3, 1}))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	i := 0
	sum := 0
	for i < len(nums) {
		sum += nums[i]
		i += 2
	}
	return sum
}
