/*
author: Rohit Sharma
Date: 18/02/2022
*/
package main

import "fmt"

/*Create function named with max */
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	var a int = 10
	var b int = 20
	var res int
	/*function call */
	res = max(a, b)
	fmt.Printf("Maximum Value is %d\n", res)
}

/* Another Example (String swapping using function) :-
package main
import "fmt"

func swap(a,b string) (string, string){
	return b,a
}
func main(){
	x,y := swap("abc","cba")
	fmt.Println(x,y)
}
*/
