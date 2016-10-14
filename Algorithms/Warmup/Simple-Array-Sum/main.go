package main

import "fmt"

func main() {
	var (
		n     int
		count int
		sum   int
	)

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &count)
		sum += count
	}
	fmt.Printf("%v\n", sum)
}
