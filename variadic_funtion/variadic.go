package main

import "fmt"

func main() {
	fmt.Println("variadic function")
	variadic(2)
	variadic(2, 3)
	b := []int{2, 3, 4}
	variadic(b...)
}

/**
Variadic functions can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function
**/

func variadic(a ...int) {
	fmt.Println("all arguments element", a)

	for k, v := range a {
		fmt.Println("element", k, v)
	}
}
