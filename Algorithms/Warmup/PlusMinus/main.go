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
		n, b                              int
		zeroVal, positiveVal, negativeVal float64
		line                              string
		allValLine                        []string
		allIntVal                         []int
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	line = scanner.Text()
	allValLine = strings.Split(line, " ")
	for i := 0; i < len(allValLine); i++ {
		b, _ = strconv.Atoi(allValLine[i])
		allIntVal = append(allIntVal, b)
	}

	for i := 0; i < n; i++ {
		switch {
		case allIntVal[i] == 0:
			zeroVal++
		case allIntVal[i] > 0:
			positiveVal++
		case allIntVal[i] < 0:
			negativeVal++
		}
	}

	fmt.Printf("%f\n", positiveVal/float64(n))
	fmt.Printf("%f\n", negativeVal/float64(n))
	fmt.Printf("%f\n", zeroVal/float64(n))
}
