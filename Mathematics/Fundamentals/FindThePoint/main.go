package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		c := make([]int, 4)
		for i := range c {
			fmt.Scanf("%d", &c[i])
		}
		fmt.Printf("%d %d\n", 2*c[2]-c[0], 2*c[3]-c[1])
	}
}
