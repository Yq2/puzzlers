package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	coordinateWithContext()
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			//如果当前goroutine的值已经是total，则向主goutine发送停止信号
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	//一共13个goutine（包含主goutine）
	//12个子routine要么成功加num后结束，要么一直循环，也就是说每个goutine最终都会成功加num一次且仅一次
	<-cxt.Done()
	fmt.Println("End.")
}

func coordinateWithContextt() {
	total := 12
	var num int32
	fmt.Println("The number :%d [with context.Context]\n", num)
	ctx , cancelFun := context.WithCancel(context.Background(),)
	for i := 0 ; i< total; i++ {
		go addNumm(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFun()
			}
		})
	}
	<- ctx.Done()
	fmt.Println("The main goutine End.")
}

// addNum 用于原子地增加一次numP所指的变量的值。
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		//获取当前循环num的值
		currNum := atomic.LoadInt32(numP)
		//然后准备在num现有值基础上加1
		newNum := currNum + 1
		//休眠停顿，200毫秒
		time.Sleep(time.Millisecond * 200)
		//如果现在num的值还是休眠之前那个值，那么就原子加1
		//为什么原子的拿到num的值，休眠很短时间之后 还要去比较交换呢
		//这是因为虽然原子的拿到num的但是拿到后并不能保证没有其他goroutine不增加num的值

		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			//如果num还是休眠之前拿到的那个值
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			//这个goroutine成功加且正确加了num的值 跳出循环，执行deferFunc函数，然后结束该goroutine
			break
		} else {
			//继续下一个循环
			//fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func addNumm(nump *int32, id int, do func()) {
	defer func() {
		do()
	}()
	for i := 0; ; i++ {
		currnum := atomic.LoadInt32(nump)
		newnum := currnum + 1
		time.Sleep(time.Millisecond * 100)
		if atomic.CompareAndSwapInt32(nump, currnum, newnum) {
			fmt.Println("The number: %d [%d-%d]\n", newnum, id,i)
			break
		} else {
			//TODO
		}
	}
}
