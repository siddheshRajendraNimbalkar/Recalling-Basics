// Reverse a string without using built-in functions.

package basic

import (
	"fmt"
	"strings"
)

func ReverseString() {
	var a string
	c := ""
	fmt.Scan(&a)
	b := strings.Split(strings.TrimSpace(a), "")
	for i := len(b) - 1; i >= 0; i-- {
		c = c + b[i]
	}
	fmt.Println(c)
}
