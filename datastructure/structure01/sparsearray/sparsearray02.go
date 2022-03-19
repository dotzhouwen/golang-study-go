package main

import "fmt"

type ValNode struct {
	Row int
	Col int
	Val int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	fmt.Println("chessMap:")
	printArray(chessMap)

	var sparseArray []ValNode

	valNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}

	sparseArray = append(sparseArray, valNode)

	for i, row := range chessMap {
		for j, v := range row {
			if v != 0 {
				sparseArray = append(sparseArray, ValNode{Row: i, Col: j, Val: v})
			}
		}
	}

	for _, s := range sparseArray {
		fmt.Printf("i:%d, j:%d, v:%v\n", s.Row, s.Col, s.Val)
	}

	var chessMap2 [11][11]int
	for i, s := range sparseArray {
		if i != 0 {
			fmt.Println(s.Row, s.Col, s.Val)
			chessMap2[s.Row][s.Col] = s.Val
		}
	}

	fmt.Println("chessMap2:")
	printArray(chessMap2)
}

func printArray(chessMap [11][11]int) {
	for _, row := range chessMap {
		for _, v := range row {
			fmt.Printf("%v  ", v)
		}
		fmt.Println()
	}
}
