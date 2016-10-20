package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	counter := make([]float32, 3)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] < 0 {
			counter[0]++
		} else if a[i] > 0 {
			counter[2]++
		} else {
			counter[1]++
		}

	}
	//fmt.Println(counter)
	formatter := "%0.3f\n"
	fmt.Printf(formatter, counter[2]/float32(n))
	fmt.Printf(formatter, counter[0]/float32(n))
	fmt.Printf(formatter, counter[1]/float32(n))
}

// func main() {
// 	var (
// 		n, b                              int
// 		zeroVal, positiveVal, negativeVal float64
// 		line                              string
// 		allValLine                        []string
// 		allIntVal                         []int
// 	)

// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ = strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	line = scanner.Text()
// 	allValLine = strings.Split(line, " ")
// 	for i := 0; i < len(allValLine); i++ {
// 		b, _ = strconv.Atoi(allValLine[i])
// 		allIntVal = append(allIntVal, b)
// 	}

// 	for i := 0; i < n; i++ {
// 		switch {
// 		case allIntVal[i] == 0:
// 			zeroVal++
// 		case allIntVal[i] > 0:
// 			positiveVal++
// 		case allIntVal[i] < 0:
// 			negativeVal++
// 		}
// 	}

// 	fmt.Printf("%f\n", positiveVal/float64(n))
// 	fmt.Printf("%f\n", negativeVal/float64(n))
// 	fmt.Printf("%f\n", zeroVal/float64(n))
// }
