package main

import (
	"fmt"
)

func main() {
	var n, a int
	var c []int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		c = append(c, a)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", c[i]+1)
	}
}
