// 4-05
package main

import (
	"fmt"
)

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for _, value := range a {
		fmt.Println(value)
	}
}
