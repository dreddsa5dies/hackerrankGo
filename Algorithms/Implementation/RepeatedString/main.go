package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		n int
		s string
	)

	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &n)

	k := n % len(s)
	a := strings.Count(s[:k], "a")
	b := strings.Count(s[k:], "a")
	res := (a+b)*(n/len(s)) + a

	fmt.Println(res)
}
