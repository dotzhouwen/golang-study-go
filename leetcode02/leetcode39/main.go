package main

import "fmt"

/**
稀疏矩阵的乘法
*/
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	row1 := len(mat1)
	col2 := len(mat2[0])
	res := make([][]int, row1)

	for i, _ := range res {
		res[i] = make([]int, col2)
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			row := getRow(&mat1, i)
			col := getCol(&mat2, j)

			val := 0
			for k := 0; k < len(row); k++ {
				val += row[k] * col[k]
			}
			res[i][j] = val
		}
	}

	return res
}

func getRow(mat *[][]int, i int) []int {
	return (*mat)[i]
}

func getCol(mat *[][]int, j int) []int {
	res := make([]int, 0)
	for _, row := range *mat {
		for i := 0; i < len(row); i++ {
			if i == j {
				res = append(res, row[i])
			}
		}
	}
	return res
}

func main() {
	mat1 := [][]int{{1, 0, 0}, {-1, 0, 3}}
	mat2 := [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	i := multiply(mat1, mat2)
	fmt.Println(i)
}
