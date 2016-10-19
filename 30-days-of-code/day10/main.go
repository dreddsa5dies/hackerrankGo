package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n      int
		counts []int
		remI   []int
	)
	scanner := bufio.NewScanner(os.Stdin)
	// scanning N
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for n > 0 {
		rem := n % 2
		n = n / 2
		remI = append(remI, rem)
	}

	count := 1
	for i := 0; i < len(remI)-1; i++ {
		if remI[i] == remI[i+1] && remI[i] == 1 {
			count++
		} else {
			count = 1
		}
		counts = append(counts, count)
	}

	b := 1
	for i := 0; i < len(counts); i++ {
		if b < counts[i] {
			b = counts[i]
		}
	}
	count = b

	if len(remI) == 1 && remI[0] == 0 {
		count = 0
	}

	fmt.Println(count)
}
