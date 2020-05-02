package main

import (
	"fmt"
)

func main() {
	var (
		n, a, d int
	)

	fmt.Scanf("%d %d", &n, &d)

	arr := []int{}

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	for _, v := range leftRotation(arr, len(arr), d) {
		fmt.Printf("%v ", v)
	}
}

func leftRotation(a []int, size int, rotation int) []int {

	var newArray []int
	for i := 0; i < rotation; i++ {
		newArray = a[1:size]
		newArray = append(newArray, a[0])
		a = newArray
	}
	return a
}
