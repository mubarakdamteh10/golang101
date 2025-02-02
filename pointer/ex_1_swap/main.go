package main

import "fmt"

func swap(a, b *int) {
	// ใส่คำตอบลงในนี้
}

func main() {
	x, y := 5, 10
	fmt.Println("Before swap:", x, y) // expect 5 10

	swap(&x, &y)

	fmt.Println("After swap:", x, y) // expect 10 5
}
