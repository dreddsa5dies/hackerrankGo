package main

import "fmt"

func main() {
	var (
		n, k, v int
	)

	fmt.Scanf("%d %d\n", &n, &k)

	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr = append(arr, v)
	}

	arrC := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if (arr[i]+arr[j])%k != 0 {
				arrC = append(arrC, arr[i])
			}
		}
	}

	counts := make(map[int]int, n)
	for i := 0; i < len(arrC); i++ {
		counts[arrC[i]] = 0
	}

	for i := 0; i < len(arrC)-1; i++ {
		if arrC[i] == arrC[i+1] {
			counts[arrC[i]]++
		}
	}

	count := 0
	for _, y := range counts {
		// не больше 0
		// а при наибольшем значении ключа!!
		if y > 0 {
			count++
		}
	}

	fmt.Println(count)
}

// DONT WORK!!!
// input
// 10 4
// 1 2 3 4 5 6 7 8 9 10

// out
// 5
