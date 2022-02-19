/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	/* Println is used for print element in next line */
	var l int64
	var u float64

	l = 289
	u = 25.30658

	fmt.Println(l)
	fmt.Println(u)

	/* Dyanamic and Mixed type decleration */
	x := 87.23
	fmt.Println(x)
	/*Print type of x varible*/
	fmt.Printf("Print Type of x  %T", x)
	fmt.Println()
	/*declare multiple variables in single line */
	var a, b, c = 2, 20.5, "ABC"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("Print Type of c  %T", c)
	fmt.Println()
}
