package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")
}

func caller() {
	fmt.Println("Enter function caller.")
	//panic()接收一个interface类型值，一般我们传入一个error类型值
	//系统运行时报告的panic异常会在向上冒泡的过程中一层一层丰富panic的异常信息
	panic(errors.New("something wrong")) // 正例。
	panic(fmt.Println)                   // 反例。
	fmt.Println("Exit function caller.")
}
