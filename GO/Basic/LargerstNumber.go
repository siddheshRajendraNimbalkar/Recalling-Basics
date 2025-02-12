// Given three numbers, determine the largest.

package basic

import (
	"fmt"
	"sort"
)

func LargestNumber() {
	fmt.Println("Enter Numbers")
	var a []int
	var b int
	for i := 0; i < 3; i++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	fmt.Println("Enter Data is: ", a)
	sort.Ints(a)
	fmt.Println("Largest data is", a[len(a)-1])
}
