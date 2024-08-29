package main

import (
	"fmt"
	"time"
)

func main() {
	todayStr := time.Now().Format("20060102")
	fmt.Println(todayStr)
	v := fmt.Sprintf("%.2f%%", float32(4)/float32(23)) // 打印百分号，100%
	fmt.Println(v)
}
