package main

import "fmt"

func main() {
	var n, k, q, a, b int

	fmt.Scanf("%d %d %d", &n, &k, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr[(i+k)%n] = a
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&b)
		fmt.Println(arr[b])
	}
}
