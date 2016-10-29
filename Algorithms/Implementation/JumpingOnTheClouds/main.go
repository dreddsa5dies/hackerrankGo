package main

import "fmt"

func main() {
	var (
		n, v int
	)

	fmt.Scanf("%d\n", &n)

	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr = append(arr, v)
	}

	count := -1
	for i := 0; i < n; i++ {
		if i < n-2 && arr[i+2] == 0 {
			i++
		}
		count++
	}

	fmt.Println(count)
}
