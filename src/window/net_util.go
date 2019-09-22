package window

import (
	"fmt"
	"github.com/dtldyg/net-speed/src/net"
)

func getDownSpeed() string {
	return fmt.Sprintf("%s/s", bytesFormat(float64(net.DownStreamDataSize)/1))
}

func getUpSpeed() string {
	return fmt.Sprintf("%s/s", bytesFormat(float64(net.UpStreamDataSize)/1))
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
