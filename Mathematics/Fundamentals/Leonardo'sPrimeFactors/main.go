package main

import (
	"fmt"
)

func main() {
	prime := [15]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	var q int
	fmt.Scanf("%d", &q)

	for i := 0; i < q; i++ {
		count := 0
		product := 1

		var n int
		fmt.Scanf("%d", &n)

		for y := 0; y < len(prime); y++ {
			product *= prime[y]
			if product <= n {
				count++
			}
		}

		fmt.Printf("%d\n", count)
	}
}
