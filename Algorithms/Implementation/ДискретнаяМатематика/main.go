package main

import "fmt"

func main() {
	var (
		l, n, k, m int
	)

	fmt.Scanf("%d", &l)
	for i := 0; i < l; i++ {
		fmt.Scanf("%d %d\n", &n, &k)
		arr := []int{}
		for j := 0; j < n; j++ {
			fmt.Scan(&m)
			arr = append(arr, m)
		}
		count := 1
		for j := 0; j < n; j++ {
			if arr[j] <= 0 {
				count++
			}
		}
		switch {
		case count <= k:
			fmt.Println("YES")
		case count > k:
			fmt.Println("NO")
		}
	}
}
