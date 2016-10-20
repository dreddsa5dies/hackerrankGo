package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n int
	)

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", n-i), strings.Repeat("#", i))
	}
}
