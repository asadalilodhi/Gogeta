package main

import (
	"fmt"
	"gogeta/iteration"
)

func main() {
	result1 := iteration.Repeat("a", 6)
	fmt.Println("Repeated:", result1)

	result2 := iteration.Repeat("b", 7)
	fmt.Println("Repeated:", result2)

	result3 := iteration.Repeat("c", 8)
	fmt.Println("Repeated:", result3)
}