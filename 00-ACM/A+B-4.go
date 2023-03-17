package main

import "fmt"

func printSumArr(arr []int) {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	fmt.Println(res)
}

func main() {
	var n, t int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			var arr []int
			for i := 0; i < n; i++ {
				fmt.Scan(&t)
				arr = append(arr, t)
			}
			printSumArr(arr)
		}
	}
}
