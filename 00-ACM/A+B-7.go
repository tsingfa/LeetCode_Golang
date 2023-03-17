package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 创建一个扫描器，从标准输入读取数据
	for scanner.Scan() {                  // 循环扫描每一行
		line := scanner.Text()           // 获取当前行的文本
		nums := strings.Split(line, " ") // 用空格分割文本，得到一个字符串切片
		sum := 0                         // 声明一个变量，用于存储和
		for _, num := range nums {       // 遍历字符串切片中的每个元素
			n, _ := strconv.Atoi(num) // 将字符串转换为整数，忽略错误信息
			sum += n                  // 累加到和中
		}
		fmt.Println(sum) // 输出和
	}
}
