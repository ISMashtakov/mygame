package actions

import (
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type Hoe struct {
	hoeHittingAnimation *images.AnimationMap
}

func NewHoe(resourcesLoader resources.IResourceLoader) *Hoe {
	return &Hoe{
		hoeHittingAnimation: resourcesLoader.LoadAnimationMap(resources.AnimationCharacterHoeHitting),
	}
}

func (h *Hoe) ProcessAction(world donburi.World, characterEntity *donburi.Entry) bool {
	if !spaceIsPressed() {
		return false
	}

	startHitAnimation(h.hoeHittingAnimation, characterEntity, func(point gmath.Vec) {
		don.Create(world, actions.GardenCreatingRequest, &actions.GardenCreatingRequestData{
			Point: point,
		})
	})

	return true
}
