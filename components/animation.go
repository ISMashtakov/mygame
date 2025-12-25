package components

import (
	"github.com/ISMashtakov/mygame/core"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
	Player core.AnimationPlayer
}

var Animation = donburi.NewComponentType[AnimationData]()

func StartAnimation(world donburi.World, player core.AnimationPlayer) *donburi.Entry {
	entity := world.Create(Animation)
	entry := world.Entry(entity)

	donburi.Add(entry, Animation, &AnimationData{Player: player})

	return entry
}

type CurrentAnimationData struct {
	Entry     *donburi.Entry
	IsWalking bool // Булевая переменная не очень мне нравится как решение, возможно на какой-то енамчик тут поменять придётся
}

var CurrentAnimation = donburi.NewComponentType[CurrentAnimationData]()
