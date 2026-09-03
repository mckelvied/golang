package main

import "fmt"

func main() {
	integers := []int{}
	for i := 1; i <= 10; i++ {
		integers = append(integers, i)
	}
	for _, value := range integers {
		if value%2 == 0 {
			fmt.Println("Integer value", value, "is even")
		} else {
			fmt.Println("Integer value", value, "is odd")
		}
	}
}
