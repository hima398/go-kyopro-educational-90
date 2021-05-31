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

	n, l := ScanInt(), ScanInt()
	k := ScanInt()
	a := ScanIntSlice(n)

	ok, ng := 0, l

	check := func(x int) bool {
		preAi := 0
		cnt := 0
		for i := 0; i < n; i++ {
			if a[i]-preAi >= x && l-a[i] >= x {
				cnt++
				preAi = a[i]
			}
		}
		return cnt >= k
	}

	for ng-ok > 1 {
		c := (ng + ok) / 2
		if check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)
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
