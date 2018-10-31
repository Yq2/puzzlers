package main

import "fmt"

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	// 不能调用不可寻址的值的指针方法。编译不报错，运行时会报错
	//New("little pig").SetName("monster")

	// 示例2。
	//第一个map没有落脚点
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
	fmt.Println("map[word]=",map1["word"]) //1


}
