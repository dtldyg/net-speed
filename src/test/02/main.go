package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"unsafe"
)

func main() {
	bitMap := robotgo.GoCaptureScreen(10, 10, 10, 10)
	//fmt.Println(bitMap.Width)
	//fmt.Println(bitMap.Height)
	//fmt.Println(bitMap.BitsPixel)
	//fmt.Println(bitMap.BytesPerPixel) // 每个像素4字节，rgba
	//fmt.Println(bitMap.Bytewidth)
	//fmt.Println(*bitMap.ImgBuf)

	fmt.Println(*bitMap.ImgBuf)
	ptr := uintptr(unsafe.Pointer(bitMap.ImgBuf))
	fmt.Println(*(*uint8)(unsafe.Pointer(ptr + unsafe.Sizeof(uint8(0)))))
	fmt.Println(*(*uint8)(unsafe.Pointer(ptr + 2*unsafe.Sizeof(uint8(0)))))
	fmt.Println(*(*uint8)(unsafe.Pointer(ptr + 3*unsafe.Sizeof(uint8(0)))))
}
