package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
将有序数组转换为二叉搜索树
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	left := 0
	right := len(nums)
	middle := (left + right) / 2

	root := &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[left:middle]),
		Right: sortedArrayToBST(nums[middle+1 : right]),
	}

	return root
}

func main() {
	nums := []int{1}
	bst := sortedArrayToBST(nums)
	fmt.Println(bst)
}
