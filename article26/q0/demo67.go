package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithWaitGroup()
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	//fmt.Println("Start loop ...")
	//外层for循环 只执行4轮
	for i := 1; i <= total; i = i + stride {
		//if i > 1 {
		//	fmt.Println("Next iteration:")
		//}
		wg.Add(stride)
		//下面的for循环会执行三次
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		//只有当前计数周期结束后才会继续后面的外层for循环
		fmt.Println("Waiting .....")
		wg.Wait()
	}
	fmt.Println("End.")
}

// addNum 用于原子地增加一次numP所指的变量的值。
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}
