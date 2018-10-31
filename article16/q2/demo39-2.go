package main

import "fmt"

func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
	//会打印0-9每个数，但是输出顺序不保证
}