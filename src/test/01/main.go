package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	fmt.Println(glfw.GetPrimaryMonitor().GetVideoMode().Height)
	fmt.Println(glfw.GetPrimaryMonitor().GetVideoMode().Width)

	glfw.WindowHint(glfw.Decorated, 0)

	window, err := glfw.CreateWindow(50, 50, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		glfw.PollEvents()
	}
}
