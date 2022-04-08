package main

import "fmt"

/**
dfs图像渲染
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	originalColor := image[sr][sc]
	dfs(&image, sr, sc, originalColor, newColor)

	return image
}

func dfs(image *[][]int, i int, j int, originalColor int, newColor int) {
	if i < 0 || i >= len(*image) || j < 0 || j >= len((*image)[0]) {
		return
	}
	if (*image)[i][j] == newColor {
		return
	}
	if (*image)[i][j] != originalColor {
		return
	}

	(*image)[i][j] = newColor
	dfs(image, i-1, j, originalColor, newColor)
	dfs(image, i+1, j, originalColor, newColor)
	dfs(image, i, j-1, originalColor, newColor)
	dfs(image, i, j+1, originalColor, newColor)
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	newColor := 2

	fill := floodFill(image, sr, sc, newColor)

	fmt.Println(fill)
}
