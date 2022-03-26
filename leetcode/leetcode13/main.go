package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	ret := removeDuplicates(nums)

	fmt.Println(nums)
	fmt.Println(ret)

}
func removeDuplicates(nums []int) int {
	slowIndex := 1
	fastIndex := 1

	for ; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] == nums[fastIndex-1] {
			continue
		}
		nums[slowIndex] = nums[fastIndex]
		slowIndex++
	}

	return slowIndex
}
