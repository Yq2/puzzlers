package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("【defer】 nter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("【defer】 panic: %s\n", p)
		}
		panic(errors.New("在defer语句里面捕获异常后，再抛出异常"))

		fmt.Println("【defer】 Exit defer function.")
	}()

	// recover函数的错误用法。运行到这里并没有发生panic异常，所以打印nil
	fmt.Printf("no panic: %v\n", recover())

	// 引发panic。
	panic(errors.New("something wrong"))

	// recover函数的错误用法。
	//下面的语句将不会执行，因为一旦发生panic异常，运行时系统会立刻沿着调用栈反方向传递，并不会继续执行
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}
