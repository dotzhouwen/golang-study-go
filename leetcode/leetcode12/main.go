package main

import "fmt"

/**
盛最多水的容器

给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49

*/

func maxArea2(height []int) int {
	left := 0
	right := len(height) - 1

	maxArea := 0
	for left < right {
		var minHeight int
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}
		area := (right - left) * minHeight
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

/**
brute force
*/
func maxArea(height []int) int {
	left := 0
	right := left + 1

	var maxArea int
	for ; right < len(height); right++ {
		for ; left < right; left++ {
			var minHeight int
			if height[left] < height[right] {
				minHeight = height[left]
			} else {
				minHeight = height[right]
			}
			area := (right - left) * minHeight
			if area > maxArea {
				maxArea = area
			}
		}
		left = 0
	}
	return maxArea
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea2(nums)
	fmt.Println("area:", area)
}
