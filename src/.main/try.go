package main

import (
	"fmt"
	"time"
)

type iterator1 struct {
	current int
	n int
}

func newIterator1(n int) *iterator1 {
	return &iterator1{0, n}
}

func (it *iterator1) hasNext() bool {
	return it.current < it.n
}

func (it *iterator1) next() int {
	if !it.hasNext() { panic("NOOOOOOOOOOOOOOOO") }
	it.current++
	return it.current - 1
}

type iterator2 <-chan int

func newIterator2(n int) iterator2 {
	result := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			result <- i
		}
		close(result)
	}()
	return result
}

func main() {
	const k = 1000000
	
	t := time.Now()
	it1 := newIterator1(k)
	i := 0
	for it1.hasNext() {
		i += it1.next()
	}
	fmt.Println(time.Now().Sub(t))
	
	t = time.Now()
	it2 := newIterator2(k)
	i = 0
	for n := range it2 {
		i += n
	}
	fmt.Println(time.Now().Sub(t))
}
