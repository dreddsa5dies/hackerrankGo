package main

import "fmt"

func main() {
	var t, n, k int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		arr := make([]int, n+1)
		if k != 0 && n%k != 0 {
			fmt.Println(-1)
		} else {
			for j := 1; j < len(arr); j++ {
				if arr[i] == 0 {
					arr[i] = i + k
					arr[i+k] = i
				}
			}
		}
		// don't work input my logic
		for j := 1; j < len(arr); j++ {
			fmt.Printf("%d", arr[j])
		}
	}
}
