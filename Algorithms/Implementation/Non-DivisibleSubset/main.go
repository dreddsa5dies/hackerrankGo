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

	arrC := map[int]int{}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				if (arr[i]%k + arr[j]%k) == k {
					arrC[arr[j]]++
				}
			}
		}
	}

	for x, _ := range arrC {
		for i := 0; i < len(arr); i++ {
			if arr[i] == x {
				arr = removeAtIndex(arr, i)
			}
		}
	}

	fmt.Println(arr)
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//меняем последнее значение и значение, которое хотим удалить, местами
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
