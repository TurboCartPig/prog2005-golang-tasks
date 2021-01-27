package main

import (
	"fmt"
	"tasks/fibandfact"
)

func main() {
	var n int
	fmt.Print("n = ")
	fmt.Scanf("%d", &n)
	fmt.Println(fibandfact.Fib(n))
}
