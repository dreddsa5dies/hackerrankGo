package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n           int
		lineMatrix  string
		allValLineM []string
		matrix      [][]string
	)
	scanner := bufio.NewScanner(os.Stdin)
	// scanning matrix of size N
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		// scanning line of matrix
		scanner.Scan()
		lineMatrix = scanner.Text()
		allValLineM = strings.Split(lineMatrix, " ")
		matrix = append(matrix, allValLineM)
	}

	fmt.Println(calculation(matrix))
}

func calculation(matrix [][]string) int {
	var (
		arrPrDiag  []int
		arrSecDiag []int
		sumPrDiag  int
		sumSecDiag int
	)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == i {
				b, _ := strconv.Atoi(matrix[i][j])
				arrPrDiag = append(arrPrDiag, b)
			}
		}
	}

	// problem in this case
	count := len(matrix[0]) - 1
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if j == count {
				b, _ := strconv.Atoi(matrix[i][j])
				arrSecDiag = append(arrSecDiag, b)
			}
		}
		count--
	}

	for i := 0; i < len(arrPrDiag); i++ {
		sumPrDiag += arrPrDiag[i]
	}

	for i := 0; i < len(arrSecDiag); i++ {
		sumSecDiag += arrSecDiag[i]
	}

	allSum := sumPrDiag - sumSecDiag

	if allSum < 0 {
		allSum = -allSum
	}
	return allSum
}
