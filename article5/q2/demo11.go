package main

import "fmt"

var container = []string{"zero1", "one1", "two1"}

func main() {
	if value, ok := interface{}(container).([]string); ok {
		fmt.Println("container is []string type, value:",value)
	}

	container := map[int]string{0: "zero2", 1: "one2", 2: "two2"}

	if value, ok := interface{}(container).(map[int]string); ok {
		fmt.Println("container is map[int]string type, value:",value)
	}

	fmt.Printf("The element is %q.\n", container[1])
}
