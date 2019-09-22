package main

import (
	"github.com/dtldyg/net-speed/src/window"
)

func main() {
	// 开启网速抓包
	//net.StartCatchSpeed()
	// 打开界面
	window.OpenWindow()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
