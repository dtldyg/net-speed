package window

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type TextSys struct {
	//TODO test entity
	t float32
}

func (sys *TextSys) New(world *ecs.World) {
	//make entity
	bgs.world = world
	bgs.bgEntity = &BGEntity{}
	// BasicEntity(id)
	bgs.bgEntity.BasicEntity = ecs.NewBasic()
	// SpaceComponent(like collider)
	bgs.bgEntity.SpaceComponent = common.SpaceComponent{}
	// RenderComponent(render)

	//add entity to system
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bgs.bgEntity.BasicEntity, &bgs.bgEntity.RenderComponent, &bgs.bgEntity.SpaceComponent)
		}
	}
}

func (sys *TextSys) Update(dt float32) {
	sys.t += dt
	if sys.t >= 1 {
		sys.t = 0
		//TODO 更新
	}
}

func (sys *TextSys) Remove(e ecs.BasicEntity) {}
