package main

import "fmt"

func main() {
	var (
		s, t, a, b, m, n, v, c int
	)

	fmt.Scanf("%d %d\n", &s, &t)
	fmt.Scanf("%d %d\n", &a, &b)
	fmt.Scanf("%d %d\n", &m, &n)

	arrApple := []int{}
	for i := 0; i < m; i++ {
		fmt.Scan(&v)
		arrApple = append(arrApple, v)
	}
	arrOrange := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		arrOrange = append(arrOrange, c)
	}

	countApple := 0
	for i := 0; i < len(arrApple); i++ {
		if (a+arrApple[i]) >= s && (a+arrApple[i]) <= t {
			countApple++
		}
	}

	countOrange := 0
	for i := 0; i < len(arrOrange); i++ {
		if (b+arrOrange[i]) >= s && (b+arrOrange[i]) <= t {
			countOrange++
		}
	}

	fmt.Println(countApple)
	fmt.Println(countOrange)
}
