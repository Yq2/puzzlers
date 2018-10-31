package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.RWMutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	//分析下sendCond条件变量获取的是读锁还是写锁
	/********/
	//sync.NewCond()参数是一个Locker接口类型
	//对于RWMutex读写锁而言，只有里面的写锁实现了Locker接口，读锁并没有直接实现这个接口
	//而是通过rlocker来代理实现了Locker接口
	//所以&lock传进去实际上是将条件变量和写锁绑定在一起
	/********/
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(lock.RLocker()) //获取的是读写锁的读锁

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 5
	go func(max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			//不要用if来检查状态，因为被唤醒的goroutine可能依然不满足执行条件
			//if mailbox == 1 {
			//	sendCond.Wait()
			//}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)
	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			//for mailbox == 0 {
			//	recvCond.Wait()
			//}
			if mailbox ==0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
