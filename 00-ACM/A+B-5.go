package main

import "fmt"

/*
func printSumArr(arr []int) {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	fmt.Println(res)
}
*/

func main() {
	var n, size, t int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&size)
		var arr []int //注意重置动态数组
		for j := 0; j < size; j++ {
			fmt.Scan(&t)
			arr = append(arr, t)
		}
		printSumArr(arr)
	}
}
