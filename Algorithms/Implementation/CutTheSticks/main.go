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
	fmt.Println(len(arr))
	cutSticks(arr)
}

func cutSticks(a []int) {
	sort.Ints(a)
	arr := []int{}
	for i := 0; i < len(a); i++ {
		if a[0] != a[i] {
			arr = append(arr, a[i]-a[0])
		}
	}
	if len(arr) > 0 {
		fmt.Println(len(arr))
		cutSticks(arr)
	}
}
