package main

import (
	"fmt"
	"strings"
)

func main() {
	// 示例1。
	var builder1 strings.Builder
	var builder2 strings.Builder
	builder2 = builder1 //直接赋值并不会影响原来的builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	builder2.WriteString("builder2 ") //不会对builder1生效
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()
	fmt.Printf("Grow the builder before len:%d \n", builder1.Len())

	// 示例2。
	fmt.Println("Grow the builder ...")
	builder1.Grow(1)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()



	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())

	fmt.Printf("The builder2 output(%d):\n\"%s\"\n", builder2.Len(), builder2.String())

}
