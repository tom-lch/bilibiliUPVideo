package main

import (
	"bilibiliSpider/internal/logic"
	"bilibiliSpider/stx"
	"fmt"
)

func main() {
	fmt.Println("这是bilibili爬虫")
	fmt.Println("打印up主列表")
	stx := stx.NewSTXContext()
	fmt.Println(stx)
	logic.GetUpInfo(stx)
}
