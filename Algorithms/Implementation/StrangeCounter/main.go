package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	rem := 3
	for t > rem {
		t = t - rem
		rem *= 2
	}

	fmt.Println(rem - t + 1)
}
