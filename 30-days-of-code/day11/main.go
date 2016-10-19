package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// count sum value hoursglass in arr
	arrH := []int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			b := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			arrH = append(arrH, b)
		}
	}

	// read max sum oh hoursglasses
	sort.Ints(arrH)

	// stdout
	fmt.Printf("%v\n", arrH[len(arrH)-1])

}
