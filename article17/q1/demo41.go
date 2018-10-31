package main

import "fmt"

func main() {
	// 示例1。
	//切片
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			//按位或
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 示例2。
	//数组
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	//这里range 的numbers2是副本
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	// 示例3。
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) - 1
	//range 的numbers3是slice ，在迭代过程中修改的是原始值
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
