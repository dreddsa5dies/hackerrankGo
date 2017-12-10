package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		for i := range c {
			fmt.Scanf("%d", &c[i])
		}
	}

	for i := range c {
		fmt.Printf("%d\n", (c[i]*(c[i]-1))/2)
	}
}
