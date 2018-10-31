package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithChan()
	fmt.Println()
	coordinateWithWaitGroup()
}

func coordinateWithChan() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)
	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})
	<-sign
	<-sign
}

func coordinateWithChann() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number :%d [with chan struct{}]\n", num)
	max := int32(10)
	go addNumm(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNumm(&num, 2, max, func() {
		sign <- struct{}{}
	})
	<- sign
	<- sign
}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)
	go addNum(&num, 3, max, wg.Done)
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
}

func coordinateWithWaitGroupp() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("the number : %d [with sync.WaitGroup]\n", num)
	max := int32(10)
	//wg.Done这种传递参数的方式相当于传递的是函数名，
	//deferFunc这是一个函数类型，五参数，无返回值
	go addNumm(&num, 2, max, wg.Done)
	go addNumm(&num, 4, max, wg.Done)
}

// addNum 用于原子地增加numP所指的变量的值。
func addNum(numP *int32, id, max int32, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		if currNum >= max {
			break
		}
		newNum := currNum + 2
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func addNumm(numP *int32, id,max int32, deferFunc func() ) {
	//deferFunc这是一个函数类型，五参数，无返回值
	defer func() {
		deferFunc()
	}()
	for i := 0 ; ; i++ {
		currNum := atomic.LoadInt32(numP)
		if currNum >= max {
			break
		}
		newNum := currNum + 2
		time.Sleep(time.Millisecond * 200)
		//currNum 是当前值，newNum是将要设置的值
		if atomic.CompareAndSwapInt32(numP,currNum,newNum) {
			fmt.Println("The number: %d [%d-%d]\n", newNum,id,i)
		} else {
			fmt.Println("The CAS operation failed .[%d-%d]\n", id, i)
		}
	}
}
