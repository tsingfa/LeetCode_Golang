/*
该系列来自
牛客：OJ在线编程常见输入输出练习场
https://ac.nowcoder.com/acm/contest/5652
用于熟悉ACM模式的输入输出
*/
package main

import "fmt"

func main() {
	var a, b int
	for {
		// 从标准输入读取两个整数，赋值给a和b，
		//返回读取的数量和错误信息
		_, err := fmt.Scan(&a, &b)
		if err != nil { //如果错误发生（即输入停止）
			break
		}
		fmt.Println(a + b)
	}

}
