package main

import "fmt"

func main() {
	var (
		n      int
		s      string
		s1, s2 string
	)

	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &s)
		s1 = ""
		s2 = ""

		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				s1 += string(s[i])
			} else {
				s2 += string(s[i])
			}
		}
		fmt.Printf("%v %v\n", s1, s2)
	}
}
