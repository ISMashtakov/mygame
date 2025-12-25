package subsystems

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type DropGeneratorSubSystem struct {
	propsCreator entities.PropsCreator
}

func NewDropGeneratorSubSystem(propsCreator entities.PropsCreator) *DropGeneratorSubSystem {
	return &DropGeneratorSubSystem{
		propsCreator: propsCreator,
	}
}

func (g DropGeneratorSubSystem) Generate(entity *donburi.Entry) {
	if entity.HasComponent(components.Position) && entity.HasComponent(components.Destroyable) {
		pos, destroyable := components.Position.Get(entity), components.Destroyable.Get(entity)

		for _, item := range destroyable.Resources {
			itemShift := gmath.Vec{X: utils.RandomFloat(-30, 30), Y: utils.RandomFloat(-30, 30)}
			prop := g.propsCreator.Create(entity.World, item, pos.Vec)
			duration := time.Millisecond * time.Duration(itemShift.Len()) * 15
			components.StartAnimation(
				entity.World,
				*core.NewAnimationPlayer(
					duration,
					core.WithOnFinish(g.GetStartIdleAnimationFunc(prop)),
					core.WithAnimations(
						animations.NewSquareMoveAnimation(itemShift, duration, components.Position.Get(prop)),
					),
				),
			)

		}
	}
}

func (g DropGeneratorSubSystem) GetStartIdleAnimationFunc(prop *donburi.Entry) func() {
	return func() {
		duration := time.Second * 3
		components.StartAnimation(
			prop.World,
			*core.NewAnimationPlayer(
				duration,
				core.WithRepetable(true),
				core.WithAnimations(
					animations.NewMoveAnimation(
						funcs.Zero(), funcs.Abs(funcs.X(duration)/2, -20), components.Position.Get(prop)),
				),
			),
		)
	}
}
