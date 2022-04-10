package main

import "fmt"

func sortColors(nums []int) {
	zero := 0
	i := 0
	two := len(nums)

	for i < two {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		}
	}
}

func main() {
	nums := []int{2}
	sortColors(nums)

	fmt.Println(nums)
}
