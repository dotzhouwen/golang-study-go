package main

import "errors"

type SqList struct {
	data    []int
	length  int
	maxsize int
}

func NewSqList(maxsize int) *SqList {
	return &SqList{
		data:    make([]int, maxsize, maxsize),
		length:  0,
		maxsize: maxsize,
	}
}

func (s *SqList) ListInsert(i int, v int) error {
	if s.length == s.maxsize {
		return errors.New("线性表已经满了")
	}
	if i < 1 || i > s.length+1 {
		return errors.New("I 不合法")
	}
	s.data[i-1] = v
	s.length++
	return nil
}
