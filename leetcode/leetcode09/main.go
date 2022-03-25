package main

import "fmt"

/**
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

*/

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		fmt.Printf("i:%d, j:%d, num[%d]=%d, num[%d]=%d, sum=%d\n", i, j, i, numbers[i], j, numbers[j], numbers[i]+numbers[j])
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] < target {
			i++
		}
	}
	return nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 19}
	target := 13
	ret := twoSum(numbers, target)

	fmt.Println(ret)
}
