package window

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type BGSys struct {
	world    *ecs.World
	bgEntity *BGEntity
}

func (bgs *BGSys) New(world *ecs.World) {
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

func (bgs *BGSys) Update(dt float32) {
	//TODO update entity.RenderComponent.Pixels
}

func (bgs *BGSys) Remove(e ecs.BasicEntity) {}

type BGEntity struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
