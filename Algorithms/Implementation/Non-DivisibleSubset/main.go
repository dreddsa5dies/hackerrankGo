package main

import "fmt"

func main() {
	var (
		n, k, v int
	)

	fmt.Scanf("%d %d\n", &n, &k)

	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		arr = append(arr, v)
	}

	rem := []int{}
	for j := 0; j < n; j++ {
		rem[j] = arr[j] % k
	}

	ss := []int{}
	ic, jc, zi, zj := 0, 0, 0, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (rem[i] + rem[j]) == k {
				ic = 0
				jc = 0
				for f := 0; f < n; f++ {
					if rem[f] == i {
						rem[f] = -1
						ic++
					} else if rem[f] == j {
						rem[f] = -1
						jc++
					}
				}
				if ic >= jc {
					for ic > 0 {
						zi = 0
						ss[zi] = i
						zi++
						ic--
					}
				} else {
					for jc > 0 {
						zj = 0
						ss[zj] = j
						zj++
						jc--
					}
				}
			}
		}
	}
	temp := n
	for r := 0; r < n; r++ {
		if rem[r] >= 0 {
			ss[temp-1] = rem[r]
		}
	}
	newcount := 0
	for nc := 0; nc < len(ss); nc++ {
		newcount++
	}

	fmt.Println(newcount)
}
