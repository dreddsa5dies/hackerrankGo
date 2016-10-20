package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		str string
	)

	fmt.Scanf("%s", &str)
	convStr := strings.Split(str, ":")
	b := 0
	if strings.HasSuffix(str, "AM") {
		if strings.HasPrefix(str, "12") {
			convStr[0] = "00"
			strP := strings.Split(str, ":")
			fmt.Printf("%s:%s:%s\n", convStr[0], strP[1], strings.Trim(strP[2], "AM"))
		} else {
			fmt.Println(strings.Trim(str, "AM"))
		}
	} else if strings.HasSuffix(str, "PM") {
		if strings.HasPrefix(str, "12") {
			fmt.Println(strings.Trim(str, "PM"))
		} else {
			b, _ = strconv.Atoi(convStr[0])
			b += 12
			strP := strings.Split(str, ":")
			fmt.Printf("%v:%s:%s\n", b, strP[1], strings.Trim(strP[2], "PM"))
		}
	} else {
		fmt.Println("12:00:00")
	}
}

// best practice
// func main() {
// 	var h, m, s, f string
// 	fmt.Scanf("%2v:%2v:%2v%2v", &h, &m, &s, &f)
// 	fmt.Println(convert(h, m, s, f))
// }
// func convert(h, m, s, f string) string {
// 	H, _ := strconv.Atoi(h)
// 	switch f {
// 	case "AM":
// 		if H == 12 {
// 			H = 0
// 		}
// 	case "PM":
// 		if H != 12 {
// 			H = H + 12
// 		}
// 	}
// 	return fmt.Sprintf("%02d:%v:%v", H, m, s)
// }
