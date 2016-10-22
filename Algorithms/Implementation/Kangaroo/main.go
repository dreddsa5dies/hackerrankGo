package main

import "fmt"

func main() {
	var x1, x2, v1, v2 int

	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)
	// example scan stdin
	// fmt.Fscanf(os.Stdin, "%d %d %d %d", &x1, &v1, &x2, &v2)
	if v2 > v1 && x2 > x1 {
		fmt.Println("NO")
	} else if v2 < v1 && x2 > x1 {
		for x2 > x1 {
			x1 = x1 + v1
			x2 = x2 + v2
		}
		switch {
		case x1 == x2:
			fmt.Println("YES")
		case x1 != x2:
			fmt.Println("NO")
		}
	} else if v2 == v1 {
		fmt.Println("NO")
	}
}
