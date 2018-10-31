package main

import (
	"context"
	"fmt"
	"time"
)

type myKey int

func main() {
	keys := []myKey{
		myKey(20),
		myKey(30),
		myKey(60),
		myKey(61),
	}
	values := []string{
		"value in node2",
		"value in node3",
		"value in node6",
		"value in node6Branch",
	}
	//context是并发安全的
	//context根节点是不带任何功能的，不带任务数据的全局唯一节点
	//可取消的context不能携带数据
	//能携带数据的context不能取消
	//取消信号是从信号起始节点往下传输
	//数据节点的value查找是从查找节点往上遍历
	//传递取消信号时，如果遇到带数据节点或忽略这个节点
	//查找数据value时，如果遇到可取消节点会忽略这个节点
	//如果在当前节点查找数据key，同时当前节点属于可取消节点那么会查找当前节点的父节点，找到就返回该值，找不到则继续往上查找
	//如果在当前节点触发取消信号，同时当前节点属于数据节点那么会忽略当前节点，并把取消信号传递给当前节点的每一个子节点
	//可撤销节点接收到撤销信号是会向它所有子值传播撤销信号，这些子值会如法炮制把撤销信号继续传播下去，最后，这个Context值会断开它与其父值之间的关联
	//携带数据节点接收到撤销信号是会向它所有子节点传播撤销信号，子节点如法炮制，但数据节点不会断开它与父节点的关联
	rootNode := context.Background()
	node1, cancelFunc1 := context.WithCancel(rootNode)
	defer cancelFunc1()

	// 示例1。
	node2 := context.WithValue(node1, keys[0], values[0])
	node3 := context.WithValue(node2, keys[1], values[1])
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[0], node3.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[1], node3.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node3: %v\n",
		keys[2], node3.Value(keys[2]))
	fmt.Println()

	// 示例2。
	node4, cancelFunc4 := context.WithCancel(node3)
	_ = cancelFunc4
	cancelFunc1()
	node5, _ := context.WithTimeout(node4, time.Hour)
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[0], node5.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node5: %v\n",
		keys[1], node5.Value(keys[1]))
	fmt.Println()

	// 示例3。
	node6 := context.WithValue(node5, keys[2], values[2])
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[0], node6.Value(keys[0]))
	fmt.Printf("The value of the key %v found in the node6: %v\n",
		keys[2], node6.Value(keys[2]))
	fmt.Println()

	// 示例4。
	node6Branch := context.WithValue(node5, keys[3], values[3])
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[1], node6Branch.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[2], node6Branch.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node6Branch: %v\n",
		keys[3], node6Branch.Value(keys[3]))
	fmt.Println()

	// 示例5。
	node7, _ := context.WithCancel(node6)
	node8, _ := context.WithTimeout(node7, time.Hour)
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[1], node8.Value(keys[1]))
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[2], node8.Value(keys[2]))
	fmt.Printf("The value of the key %v found in the node8: %v\n",
		keys[3], node8.Value(keys[3]))
}
