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
	var (
		n, m, answer, result int
	)
	fmt.Scanf("%d %d", &n, &m)
	arrCityes := make([]int, m)
	if n != m {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		arrStations := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(arrStations); i++ {
			st, _ := strconv.Atoi(arrStations[i])
			arrCityes[i] = st
		}
		sort.Ints(arrCityes)
		if arrCityes[0] != 0 {
			answer = arrCityes[0]
		}
		if arrCityes[m-1] != n-1 {
			result = (n - 1) - arrCityes[m-1]
		}
		if result > answer {
			answer = result
		}

		for i := 0; i < m-1; i++ {
			if (arrCityes[i+1]-arrCityes[i])%2 == 0 {
				result = (arrCityes[i+1] - arrCityes[i]) / 2
			} else {
				result = ((arrCityes[i+1] - arrCityes[i]) - 1) / 2
			}
			if result > answer {
				answer = result
			}
		}

		fmt.Println(answer)
	} else {
		fmt.Println(0)
	}

}
