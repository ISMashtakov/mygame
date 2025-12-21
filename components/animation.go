package components

import (
	"github.com/ISMashtakov/mygame/core"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
	Player core.AnimationPlayer
}

var Animation = donburi.NewComponentType[AnimationData]()

func StartAnimation(world donburi.World, player core.AnimationPlayer) {
	entity := world.Create(Animation)
	entry := world.Entry(entity)

	donburi.Add(entry, Animation, &AnimationData{Player: player})
}

func DeleteAnimation(world donburi.World, an *donburi.Entry) {
	world.Remove(an.Entity())
}
