package main

import "fmt"

/**
00,  01,  02

10,  11,  12

20,  21,  22





*/

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}

	for i := 0; i < length/2; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j], matrix[length-i-1][j] = matrix[length-i-1][j], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 5},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	rotate(matrix)

	for _, row := range matrix {
		fmt.Println(row)
	}
}
