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
	arrHour := [][]int{}

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

	// scans hourglass on arr
	hourglass := []int{}
	for i := 0; i < 3; i++ {
		switch {
		case i == 0 || i == 2:
			for j := 0; j < 3; j++ {
				hourglass = append(hourglass, arr[i][j])
			}
		case i == 1:
			hourglass = append(hourglass, arr[i][1])
		}
	}
	arrHour = append(arrHour, hourglass)

	// stdout
	fmt.Printf("%v ", arrHour)

}
