package main

import (
	"container/ring"
	"fmt"
)
//ring可以用来保存固定数量的元素，例如保存最近100条日志，用户最近10次操作

func main() {
	var ringitem *ring.Ring
	//ring默认是一个长度为1的循环链表
	fmt.Println("ringitem Len:",ringitem.Len())
	var ringitem2 *ring.Ring = ring.New(5)
	fmt.Println("ringitem2 Len:",ringitem2.Len())

	//item := ringitem
	//for {
	//	if item.Value == nil {
	//		break
	//	}
	//	fmt.Println("item:",item.Value)
	//	item = item.Next()
	//}
	ringitem2.Value = 3
	ringitem2.Next().Value = 4
	ringitem2.Next().Next().Value = 5
	item2 := ringitem2
	for {
		if item2.Value == nil {
			break
		}
		fmt.Println("ringitem2:",item2.Value)
		item2 = item2.Next()
	}
	//ringitem2.Move(-1)
	//ringitem2.Unlink(5)
	ringitem2.Do(func(i interface{}) {
		//i = 8
		fmt.Println("i:",i)
	})
}
