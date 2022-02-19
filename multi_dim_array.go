/*
author: Rohit Sharma
Date: 19/02/2022
*/
package main

import "fmt"

func main() {
	var a = [5][3]int{{0, 10, 100}, {1, 20, 200}, {2, 30, 300}}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
