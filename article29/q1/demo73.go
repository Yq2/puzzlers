package main

import (
	"fmt"
)

func main() {
	str := "Go爱好者"
	//不同类型的字符占用的字节是不一样的
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
