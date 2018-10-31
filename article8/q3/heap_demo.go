package main

import "container/heap"

//heap可以用来排序。游戏编程中是一种高效的定时器实现方案
type myHeap struct {
	number int
	x int
	y int
	list []interface{}
}

func (p *myHeap) Len() int {
	return p.number
}

func (p *myHeap) Less(i,j int) bool {
	return p.x > p.y
}

func (p *myHeap) Swap(i, j int)  {
	p.x, p.y = p.y, p.x
}

func (p *myHeap) Push(x interface{}) {
	p.list = append(p.list,x)
	p.number++
}

func (p *myHeap) Pop() interface{} {
	pop := p.list[len(p.list)-1]
	p.list = p.list[:len(p.list)-1]
	return pop
}

func main() {
	var myheap = &myHeap{}
	heap.Init(myheap)
}
