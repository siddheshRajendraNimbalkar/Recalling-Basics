// Swap two variables without using a third variable.
package basic

import "fmt"

func Swap() {
	var a int = 1
	var b int = 2
	fmt.Println("a:", a, " b:", b)
	a, b = b, a
	fmt.Println("a:", a, " b:", b)
}
