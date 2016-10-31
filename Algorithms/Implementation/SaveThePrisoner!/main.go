package main

import "fmt"

func main() {
	var t, n, m, s int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %d %d", &n, &m, &s)
		x := (s + m - 1) % n
		if x == 0 {
			fmt.Println(n)
		} else {
			fmt.Println(x)
		}
	}
}
