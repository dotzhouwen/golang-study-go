package main

func RescuiveTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuiveTail(n - 1, a * n)
}


