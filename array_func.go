/*
author: Rohit Sharma
Date: 19/02/2022
*/
package main

import "fmt"

func getavg(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}

func main() {
	var b = []int{200, 3500, 70, 15, 5}
	var avg float32
	avg = getavg(b, 5)
	fmt.Printf("Average is %f", avg)
}
