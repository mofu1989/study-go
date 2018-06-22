package main

import (
	"github.com/otiai10/gosseract"
	//"fmt"
	"fmt"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("D:\\liuyonghui1\\桌面\\工作\\图片识别\\test.png")
	text, _:=client.Text()
	fmt.Println(text)
}
