package main

import (
	"fmt"
	//"time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
	//打印的结果是不确定的，可能全部是10也可能部分是10，结果中并不一定包含0-9中的每个数字
}
