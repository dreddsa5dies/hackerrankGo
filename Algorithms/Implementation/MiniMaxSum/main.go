package main

import (
	"fmt"
)

func main() {
	var (
		a, b, x []int
		c, d, v int
	)
	for i := 0; i < 5; i++ {
		fmt.Scan(&v)
		x = append(x, v)
	}
	length := len(x)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	a = x[1:]
	b = x[:4]
	for i := 0; i < 4; i++ {
		c += a[i]
		d += b[i]
	}
	fmt.Println(d, c)
}
