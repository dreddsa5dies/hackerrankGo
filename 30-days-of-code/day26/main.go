package main

import "fmt"

func main() {
	var (
		d1, d2, m1, m2, y1, y2 int
	)
	fmt.Scanln(&d1, &m1, &y1)
	fmt.Scanln(&d2, &m2, &y2)

	fine := 0

	if y1 < y2 {
		fine = 0
	} else {
		if y1 > y2 {
			fine = 10000
		} else if m1 > m2 {
			fine = 500 * (m1 - m2)
		} else if d1 > d2 {
			fine = 15 * (d1 - d2)
		}
	}

	fmt.Println(fine)
}
