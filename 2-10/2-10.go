// 2-10
package main

import (
	"fmt"
)

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Thursday)
	fmt.Println(numberOfDays)
}
