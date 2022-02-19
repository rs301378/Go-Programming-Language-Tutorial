/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	var x string = "B"
	switch {
	case x == "A":
		fmt.Print("Grade A")
	case x == "B":
		fmt.Print("Grade B")
	case x == "C":
		fmt.Print("Grade C")
	case x == "D":
		fmt.Print("Grade D")
	default:
		fmt.Print("No Grade")
	}
}
