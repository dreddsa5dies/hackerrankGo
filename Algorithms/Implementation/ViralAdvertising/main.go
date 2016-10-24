package main

import "fmt"

func main() {
	var n, x, s int

	fmt.Scanf("%d", &n) // denoting a number of days
	count := 5
	for i := 0; i < n; i++ {
		x = (count / 2) * 3
		s = s + (count / 2)
		count = x
	}
	fmt.Println(s)
}
