package main

import (
	"github.com/dtldyg/net-speed/src/net"
	"github.com/dtldyg/net-speed/src/window"
)

func main() {
	// 开启网速抓包
	go net.StartCatchSpeed()
	// 打开界面
	window.OpenWindow()
}
