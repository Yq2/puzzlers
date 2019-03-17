package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("note",os.O_RDWR,0666)
	if err != nil {
		fmt.Printf("file open file error: %v",err)
	}
	fmt.Printf("file content:%s",file)
}
