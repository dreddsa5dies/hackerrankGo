package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := []string{}
	arr := [][]int{}

	// stdin add in 2d arr
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 6; i++ {
		scanner.Scan()
		s = strings.Split(scanner.Text(), " ")
		row := []int{}
		for j := 0; j < 6; j++ {
			b, _ := strconv.Atoi(s[j])
			row = append(row, b)
		}
		arr = append(arr, row)
	}

	// stdout 2d arr
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
