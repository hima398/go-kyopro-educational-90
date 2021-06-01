package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, p, q := ScanInt(), ScanInt(), ScanInt()
	a := ScanIntSlice(n)
	ans := 0
	for i := 0; i < n-4; i++ {
		for j := i + 1; j < n-3; j++ {
			for k := j + 1; k < n-2; k++ {
				for l := k + 1; l < n-1; l++ {
					for m := l + 1; m < n; m++ {
						pi := a[i] * a[j]
						pi %= p
						pi *= a[k]
						pi %= p
						pi *= a[l]
						pi %= p
						pi *= a[m]
						pi %= p
						if pi == q {
							ans++
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func ScanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func ScanIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = ScanInt()
	}
	return s
}
