// Write a function to check if a number is even or odd.
package basic

import "fmt"

func EvenOdd() {
	fmt.Println("Enter data for Even and Odd Verification")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Enter data is Even")
	} else {
		fmt.Println("Enter data is Odd")
	}
}
