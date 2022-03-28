package main

import (
	"fmt"
	"sort"
)

/**

只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？


*/
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			i += 2
			continue
		} else {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func singleNumber2(nums []int) int {
	reduce := 0
	for _, num := range nums {
		reduce = reduce ^ num
	}
	return reduce
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	ret := singleNumber2(nums)
	fmt.Println(ret)

	fmt.Println("-------------------")
	fmt.Println(10 ^ 0)
	fmt.Println(1000 ^ 0)
	fmt.Println(10000 ^ 10000)
	fmt.Println(100000 ^ 2 ^ 100000 ^ 2)
}
