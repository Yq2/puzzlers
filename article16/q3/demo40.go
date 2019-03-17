package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var number uint32 = 10
	//  count相当于一个接力棒
	var count uint32
	trigger := func(i uint32, fn func()) {
		//  自旋锁
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				//  一定要在执行完函数后才原子加1
				atomic.AddUint32(&count, 1)
				// 只有count的值增加1后才轮到其他go执行
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
	}

	for i := uint32(0); i < number; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	//把最后一个计数交给
	trigger(number, func() {})
	//会按照自然数顺序打印（一定是这样）
}
