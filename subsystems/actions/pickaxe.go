package actions

import (
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type Pickaxe struct {
	hoeHittingAnimation *images.AnimationMap
}

func NewPickaxe(resourcesLoader resources.IResourceLoader) *Pickaxe {
	return &Pickaxe{
		hoeHittingAnimation: resourcesLoader.LoadAnimationMap(resources.AnimationCharacterHoeHitting),
	}
}

func (p *Pickaxe) ProcessAction(world donburi.World, characterEntity *donburi.Entry) bool {
	if !spaceIsPressed() {
		return false
	}

	startHitAnimation(p.hoeHittingAnimation, characterEntity, func(point gmath.Vec) {
		don.Create(world, actions.PickaxeHitRequest, &actions.PickaxeHitRequestData{
			Point: point,
		})
	})

	return true
}
