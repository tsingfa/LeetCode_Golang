package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var str []string
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		str = append(str, s)
	}
	sort.Strings(str)
	for i := 0; i < n-1; i++ {
		fmt.Printf("%v ", str[i])
	}
	fmt.Printf("%v", str[n-1])

}
