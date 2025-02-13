// Implement a function to compute the factorial of a number using recursion.

package basic

import "fmt"

func fact(num int) int {
	if num <= 0 {
		return 1
	}
	return num * fact(num-1)
}

func Recursion() {
	var num int
	fmt.Println("Enter the Data for recursion")
	fmt.Scan(&num)
	total := fact(num)
	fmt.Println(total)
}
