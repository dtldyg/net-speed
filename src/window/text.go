package window

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"image/color"
)

type TextSys struct {
	text Text
	fnt  *common.Font
	t    float32
}

func (sys *TextSys) New(world *ecs.World) {
	sys.fnt = &common.Font{
		URL:  "Consola.ttf",
		FG:   color.Black,
		Size: 10,
	}
	err := sys.fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	sys.text = Text{BasicEntity: ecs.NewBasic()}
	sys.text.RenderComponent.Drawable = common.Text{
		Font: sys.fnt,
		Text: "net\nspeed",
	}
	sys.text.SetShader(common.TextShader)
	sys.text.RenderComponent.SetZIndex(1)
	sys.text.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 5, Y: 0},
		Width:    width,
		Height:   height,
	}

	//add entity to system
	for _, system := range world.Systems() {
		switch s := system.(type) {
		case *common.RenderSystem:
			s.Add(&sys.text.BasicEntity, &sys.text.RenderComponent, &sys.text.SpaceComponent)
		}
	}
}

func (sys *TextSys) Update(dt float32) {
	sys.t += dt
	if sys.t >= 1 {
		sys.text.RenderComponent.Drawable = common.Text{
			Font: sys.fnt,
			Text: fmt.Sprintf("Down:%s\nUp:%s", getDownSpeed(float64(sys.t)), getUpSpeed(float64(sys.t))),
		}
		sys.t = 0
	}
}

func (sys *TextSys) Remove(e ecs.BasicEntity) {}

type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
