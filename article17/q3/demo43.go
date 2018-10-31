package main

import "fmt"

func main() {
	// 示例1。
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] { // 这条语句无法编译通过。
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	//case中的子表达式不能重复，只是针对字面量常量
	// 示例2。
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}

	// 示例3。
	//value6 := interface{}(byte(127))
	//switch t := value6.(type) { // 这条语句无法编译通过。
	//case uint8, uint16:
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Printf("byte")
	//default:
	//	fmt.Printf("unsupported type: %T", t)
	//}


	//无法通过编译
	value7 := [...]int8{0,1,2,3,4,5,6}
	switch value7[4] {
	case int8(0),int8(1),int8(2):
		fmt.Println("0 or 1 or 2")
	case int8(2),int8(3),int8(4):
		fmt.Println("2 or 3 or 4")
	case int8(4),int8(5),int8(6):
		fmt.Println("4 or 5 or 6")
	}
}
