/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

func main() {
	/* we cannot change constant value. Once declared, it is fixed. */
	const k int = 50

	fmt.Println("Constant value of k is: ", k)

	/* Special Sequence: are used in between the string for decoration/formatting the string.
	\n is used for next line
	\t is used for tab
	*/
	fmt.Println("When we will go \n for a walk \t next time?")
}
