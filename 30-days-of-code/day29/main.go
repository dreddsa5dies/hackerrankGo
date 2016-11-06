package main

import "fmt"

func main() {
	var (
		t, n, k int
	)
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		if ((k - 1) | k) <= n {
			fmt.Println(k - 1)
		} else {
			fmt.Println(k - 2)
		}
	}
}
