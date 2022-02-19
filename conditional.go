/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	var x int = 92
	if x >= 90 {
		fmt.Print("A \n")
	} else if x >= 80 {
		fmt.Print("B \n")
	} else if x >= 50 {
		fmt.Print("C \n")
	} else {
		fmt.Print("Default value ")
	}
}
