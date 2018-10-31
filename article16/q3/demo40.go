package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var number uint32 = 100
	//count相当于一个接力棒
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
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
	trigger(number, func(){})
	//会按照自然数顺序打印（一定是这样）
}



