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
//     var h, m, s int
//     var ap string
//     fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &ap)
//     if ap == "PM" {
//         if h < 12 { h += 12 }
//     } else if h == 12 {
//         h = 0
//     }
//     fmt.Printf("%02d:%02d:%02d", h%24, m, s)
// }
