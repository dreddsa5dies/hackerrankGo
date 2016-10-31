package main

import "fmt"

func main() {
	var (
		n  int
		nN int
		nn []int
	)

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nN)
		nn = append(nn, nN)
	}

	length := len(nn)
	count := 0
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if nn[j] > nn[j+1] {
				nn[j], nn[j+1] = nn[j+1], nn[j]
				count++
			}
		}
	}
	fmt.Printf("Array is sorted in %v swaps.\n", count)
	fmt.Printf("First Element: %v\n", nn[0])
	fmt.Printf("Last Element: %v\n", nn[len(nn)-1])
}
