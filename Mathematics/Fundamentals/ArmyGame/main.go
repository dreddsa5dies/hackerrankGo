package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	fmt.Printf("%d\n", (n+n%2)*(m+m%2)/4)
}
