package main

import (
	"fmt"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"image/color"
	"os"
	"time"

	//_ "github.com/EngoEngine/ecs"
	"github.com/google/gopacket/pcap"
	"net"
)

const (
	deviceName = "en0"
)

var (
	downStreamDataSize = 0 // 单位时间内下行的总字节数
	upStreamDataSize   = 0 // 单位时间内上行的总字节数
)

func main() {
	ifs, err := pcap.FindAllDevs()
	checkErr(err)
	var device pcap.Interface
	for _, d := range ifs {
		if d.Name == deviceName {
			device = d
		}
	}
	ipv4 := findDeviceIpv4(device)
	macAddr := findMacAddrByIp(ipv4)
	fmt.Println("ipv4:", ipv4)
	fmt.Println("mac:", macAddr)

	// 获取网卡handler，可用于读取或写入数据包
	handle, err := pcap.OpenLive(deviceName, 1024 /*每个数据包读取的最大值*/, true /*是否开启混杂模式*/, 30*time.Second /*读包超时时长*/)
	if err != nil {
		//panic(err)
	}
	defer handle.Close()

	//// 开启子线程，每一秒计算一次该秒内的数据包大小平均值，并将下载、上传总量置零
	//go output()
	//// 开始抓包
	//go monitor(handle, macAddr)

	// 打开界面
	openEngo()
}

type WindowScene struct {
}

func (ws *WindowScene) Preload() {
}

func (ws *WindowScene) Setup(u engo.Updater) {
	//set default bg
	common.SetBackground(color.White)
}

func (ws *WindowScene) Type() string {
	return "WindowScene"
}

func openEngo() {
	opts := engo.RunOptions{
		Title:          "Hello Engo!",
		Width:          60,
		Height:         60,
		StandardInputs: true,
		NotResizable:   true,
	}
	scene := &WindowScene{}
	err := glfw.Init()
	checkErr(err)
	glfw.WindowHint(glfw.Decorated, 0) // 关闭标题栏及边框等
	engo.Run(opts, scene)
}

func monitor(handle *pcap.Handle, macAddr string) {
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// 只获取以太网帧
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernet := ethernetLayer.(*layers.Ethernet)
			// 如果封包的目的MAC是本机则表示是下行的数据包，否则为上行
			if ethernet.DstMAC.String() == macAddr {
				downStreamDataSize += len(packet.Data()) // 统计下行封包总大小
			} else {
				upStreamDataSize += len(packet.Data()) // 统计上行封包总大小
			}
		}
	}
}

// 每一秒计算一次该秒内的数据包大小平均值，并将下载、上传总量置零
func output() {
	for {
		//\r:光标回到本行开头
		_, _ = os.Stdout.WriteString(fmt.Sprintf("\rDown:%s/s\tUp:%s/s", bytesFormat(float64(downStreamDataSize)/1), bytesFormat(float64(upStreamDataSize)/1)))
		downStreamDataSize = 0
		upStreamDataSize = 0
		time.Sleep(1 * time.Second)
	}
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

// 获取网卡的IPv4地址
func findDeviceIpv4(device pcap.Interface) string {
	for _, addr := range device.Addresses {
		if ipv4 := addr.IP.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	panic("device has no IPv4")
}

// 根据网卡的IPv4地址获取MAC地址
// 有此方法是因为gopacket内部未封装获取MAC地址的方法，所以这里通过找到IPv4地址相同的网卡来寻找MAC地址
func findMacAddrByIp(ip string) (string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(interfaces)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			if a, ok := addr.(*net.IPNet); ok {
				if ip == a.IP.String() {
					return i.HardwareAddr.String()
				}
			}
		}
	}
	panic(fmt.Sprintf("no device has given ip: %s", ip))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
