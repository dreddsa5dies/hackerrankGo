package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n []int
	var noS string

	// scan 1 element (not need)
	fmt.Scanf("%s\n", &noS)

	// scan array with string & append to slice
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Split(scanner.Text(), " ")
	for _, v := range str {
		val, _ := strconv.Atoi(v)
		n = append(n, val)
	}

	// reverse slice
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}

	// stdout
	for i, v := range n {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
