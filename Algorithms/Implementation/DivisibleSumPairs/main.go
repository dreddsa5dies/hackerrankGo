package main

import "fmt"

func main() {
	var (
		n, k, x, count int
	)

	fmt.Scanf("%d %d", &n, &k)
	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (arr[i]+arr[j])%k == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
