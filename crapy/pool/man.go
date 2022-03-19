package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	exists := FileExists("D:\\video\\21\\000000001s.ts")
	fmt.Println(exists)
}

func RunTest(i int, w *sync.WaitGroup) {
	fmt.Println(i)
	w.Done()
}

func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}