package main

import "fmt"

func main() {
	var (
		n, q, a int
	)
	pages, ans := 0, 0

	fmt.Scanln(&n, &q)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		times := a / q
		remainder := a % q
		questions := 0
		for j := 0; j < times; j++ {
			pages = pages + 1
			first := questions + 1
			last := questions + q
			questions = questions + q
			if pages >= first && pages <= last {
				ans = ans + 1
			}
		}
		if remainder != 0 {
			pages = pages + 1
			first := questions + 1
			last := questions + remainder
			if pages >= first && pages <= last {
				ans = ans + 1
			}
		}
	}

	fmt.Println(ans)
}
