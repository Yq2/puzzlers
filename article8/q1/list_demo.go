package main

import (
	"container/list"
	"fmt"
)
//list可以作为queue和stack的基础数据结构


func main() {
	//链表是开箱即用的（内部使用了延迟初始化）
	//指针相等是链表已经初始化的充分必要条件
	//list的零值是长度为0的空链表
	//list不是并发安全
	var listItems list.List
	//List中的根元素不会持有实际元素值
	fmt.Println("listitems Len:",listItems.Len())
	listItems.Init() //不初始化也行
	var listItems2 list.List
	fmt.Println("add 1 to listitems Front.")
	listItems.PushFront(1)
	fmt.Printf("listitems front: %v\n",listItems.Front().Value)
	fmt.Println("add 2 to listitems Back.")
	listItems.PushBack(2)
	fmt.Printf("listitems back: %v\n",listItems.Back().Value)
	fmt.Println("add 3 to listitems Back.")
	listItems.PushBack(3)
	fmt.Printf("listitems back: %v\n",listItems.Back().Value)
	fmt.Println("Remove listitems Back item .")
	//listItems.Remove(listItems.Back())
	//fmt.Printf("Removed listitems back: %v\n",listItems.Back().Value)
	//fmt.Printf("listitems Len: %d\n",listItems.Len())
	fmt.Println("add 5 to listitems Back")
	listItems.PushBack(5)
	fmt.Println("Before 5 add 4.")
	listItems.InsertBefore(4,listItems.Back())

	fmt.Println("After 5 add 6.")
	listItems.InsertAfter(6,listItems.Back())

	fmt.Println("After 6 add 8.")
	listItems.InsertAfter(8,listItems.Back())

	fmt.Println("After 8 add 7.")
	listItems.InsertAfter(7,listItems.Back())
	fmt.Println("没有排序之前：")
	item := listItems.Front()
	for {
		if item == nil {
			break
		}
		fmt.Println("listitems item:",item.Value)
		item = item.Next()
	}

	listItems.MoveAfter(listItems.Back().Prev(),listItems.Back())
	fmt.Println("排序之后:")
	item = listItems.Front()
	for {
		if item == nil {
			break
		}
		fmt.Println("listitems item:",item.Value)
		item = item.Next()
	}
	fmt.Println("listitem2.")
	listItems2.PushBack(9)
	listItems2.PushBack(10)
	listItems2.PushBack(11)
	listItems2.PushBack(12)
	item2 := listItems2.Front()
	for {
		if item2 == nil {
			break
		}
		fmt.Println("listitems2 item:",item2.Value)
		item2 = item2.Next()
	}
	fmt.Println("Push listitem2 to listitem Back.")
	listItems.PushBackList(&listItems2)
	fmt.Println("push listitem2 to listitem back after listitem .")
	item = listItems.Front()
	for {
		if item == nil {
			break
		}
		fmt.Println("listitems item:",item.Value)
		item = item.Next()
	}
}
