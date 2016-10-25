package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n, a int
	)

	fmt.Scanf("%d", &n)
	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	sort.Ints(arr)

	t := 0
	for i := 0; i < n-1; i++ {
		// если i==i+1, то проскакиваем к i+2
		if arr[i] == arr[i+1] {
			t++
			i++
		}
	}
	fmt.Println(t)
}
