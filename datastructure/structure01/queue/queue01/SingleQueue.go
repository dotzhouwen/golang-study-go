package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array [4]int
	front int
	rear int
}

func (q *Queue) Push(val int) (err error) {
	// 先判断队列是否已经满了
	if q.rear == q.maxSize - 1 {
		return errors.New("queue is full")
	}
	q.rear++
	q.array[q.rear] = val
	return nil
}

func (q *Queue) ShowQueue() {
	for i := q.front + 1; i <= q.rear; i++ {
		v := q.array[i]
		fmt.Printf("arr[%d]=%d\t", i, v)
	}
	fmt.Println()
}

func (q *Queue) GetQueue() (val int, err error) {
	if q.front == q.rear {
		return 0, errors.New("queue is empty")
	}
	q.front++
	val = q.array[q.front]
	return val, nil
}

func main() {
	queue := &Queue{maxSize: 4, front: -1, rear: -1}

	var key string
	var val int

	for {
		fmt.Println("1. 输入add表示添加数据到队列\n2. 输入get表示获取数据\n3. 输入show表示显示队列\n4.输入exit表示退出")
		_, err := fmt.Scanln(&key)
		if err != nil {
			return
		}
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数字:")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("取出数据:%v\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
