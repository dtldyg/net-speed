package window

import (
	"fmt"
	"github.com/dtldyg/net-speed/src/net"
)

func getDownSpeed(t float64) string {
	s := fmt.Sprintf("%s/s", bytesFormat(float64(net.DownStreamDataSize)/t))
	net.DownStreamDataSize = 0
	return s
}

func getUpSpeed(t float64) string {
	s := fmt.Sprintf("%s/s", bytesFormat(float64(net.UpStreamDataSize)/t))
	net.UpStreamDataSize = 0
	return s
}

// 字节数格式化
func bytesFormat(b float64) string {
	if b < 1000 {
		return fmt.Sprintf("%.1fKB", b)
	} else if b < 1024*1000 {
		return fmt.Sprintf("%.1fMB", b/1024)
	} else {
		return fmt.Sprintf("%.1fGB", b/1024/1024)
	}
}
