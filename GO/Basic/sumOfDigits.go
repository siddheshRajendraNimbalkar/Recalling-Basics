// Write a function that calculates the sum of digits of a given number.

package basic

import (
	"fmt"
	"strconv"
	"strings"
)

func SumOfDigits() {
	var num int

	fmt.Scan(&num)

	num1 := strings.Split(strconv.Itoa(num), "")
	total := 0
	for i := 0; i < len(num1); i++ {
		digit, err := strconv.Atoi(num1[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		total = total + digit
	}

	fmt.Println(total)
}
