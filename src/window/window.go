package window

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/go-gl/glfw/v3.2/glfw"
	"image/color"
)

const (
	width  = 90
	height = 30
)

func OpenWindow() {
	opts := engo.RunOptions{
		Title:          "Net-Speed",
		Width:          width,
		Height:         height,
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
	err := engo.Files.Load("Consola.ttf")
	if err != nil {
		panic(err)
	}
}

func (ws *WindowScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&TextSys{})
	common.SetBackground(color.White)
	x := 0
	y := glfw.GetPrimaryMonitor().GetVideoMode().Height - height
	engo.Window.SetPos(x, y)
}

func (ws *WindowScene) Type() string {
	return "WindowScene"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
