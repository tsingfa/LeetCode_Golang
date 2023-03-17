package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //读取所有的系统输入
	for scanner.Scan() {                  //循环扫描每一行
		line := scanner.Text()           //对于本行
		strs := strings.Split(line, ",") //将每行字符串以空格分开，存于数组
		sort.Strings(strs)
		n := len(strs)
		for i := 0; i < n-1; i++ {
			fmt.Printf("%v,", strs[i])
		}
		fmt.Printf("%v\n", strs[n-1])
	}
}
