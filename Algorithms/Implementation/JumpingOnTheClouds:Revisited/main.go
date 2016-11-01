package main

import "fmt"

func main() {
	var n, k, e int
	fmt.Scanf("%d %d", &n, &k)
	c := make([]int, n)
	for i := range c {
		fmt.Scanf("%d", &c[i])
	}

	e = 100 - n/k
	for i := 0; i < n; i = i + k {
		if c[i] == 1 {
			e = e - 2
		}
	}
	fmt.Println(e)
}
