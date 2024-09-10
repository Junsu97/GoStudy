// 4-08
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 10)

	fmt.Println(a[4])
	fmt.Println(a[5]) // 길이를 벗어난 인덱스 접근
	fmt.Println(a[8]) // 길이를 벗어난 인덱스 접근
}
