/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	var a int = 10
	for a < 30 {
		if a == 15 {
			a++
			continue
		}
		fmt.Printf("Value is %d\n", a)
		a++
		if a > 20 {
			break
		}
	}
}
