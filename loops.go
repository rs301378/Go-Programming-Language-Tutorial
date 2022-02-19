/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	var a int = 15
	var b int = 0

	for c := 0; c < 10; c++ {
		fmt.Printf("Value of a is: %d\n", c)
	}

	for b < a {
		b++
		fmt.Printf("Value of b is: %d\n", b)
	}
}
