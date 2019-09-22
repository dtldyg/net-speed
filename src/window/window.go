package window

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/go-gl/glfw/v3.2/glfw"
	"image/color"
)

func OpenWindow() {
	opts := engo.RunOptions{
		Title:          "Net-Speed",
		Width:          120,
		Height:         30,
		StandardInputs: true,
		NotResizable:   true,
	}
	scene := &WindowScene{}
	err := glfw.Init()
	checkErr(err)
	glfw.WindowHint(glfw.Decorated, 0) // 关闭标题栏及边框等
	engo.Run(opts, scene)
}

type WindowScene struct {
}

func (ws *WindowScene) Preload() {
}

func (ws *WindowScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	common.SetBackground(color.RGBA64{R: 0xffff, G: 0xffff, B: 0xffff, A: 0xffff})
	//bitMap := robotgo.GoCaptureScreen() //x,y,w,h
	world.AddSystem(&common.RenderSystem{})
}

func (ws *WindowScene) Type() string {
	return "WindowScene"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
