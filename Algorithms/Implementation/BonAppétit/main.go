package main

import "fmt"

func main() {
	var n, k, b int
	fmt.Scanf("%d %d", &n, &k)
	c := make([]int, n)
	for i := range c {
		fmt.Scanf("%d", &c[i])
	}
	fmt.Scanf("%d", &b)
	sum := 0
	for i := range c {
		if i == k {
			continue
		}
		sum += c[i]
	}
	share := sum / 2
	if share == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - share)
	}
}
