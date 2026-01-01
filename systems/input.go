package systems

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/subsystems/actions"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	Speed         = 2
	InputCodename = "input"
)

type Input struct {
	core.BaseSystem

	walkingAnimationMap *images.AnimationMap
	actionsProcessor    *actions.ActionsProcessor
}

func NewInput(resourcesLoader resources.IResourceLoader) *Input {
	return &Input{
		BaseSystem: core.BaseSystem{
			Codename:    InputCodename,
			NextSystems: []string{CollissionDetectorCodename, MovementCodename},
		},
		actionsProcessor:    actions.NewActionProcessor(resourcesLoader),
		walkingAnimationMap: resourcesLoader.LoadAnimationMap(resources.AnimationCharacterWalking),
	}
}

func (m *Input) Update(world donburi.World) {
	characterEntity, ok := donburi.NewQuery(filter.Contains(components.Character)).First(world)
	if !ok {
		panic("can't found character")
	}

	anim := components.CurrentAnimation.Get(characterEntity)
	if anim.Entry != nil && !anim.IsWalking {
		return
	}

	justPressedKeys := inpututil.AppendJustPressedKeys(nil)
	keys := inpututil.AppendPressedKeys(nil)

	if lo.Contains(justPressedKeys, ebiten.KeyTab) {
		don.Create[any](world, gui.SwitchInventaryStatusRequest, nil)
	}

	m.processNumbers(characterEntity, justPressedKeys)

	if m.actionsProcessor.Process(characterEntity) {
		return
	}

	m.processMoving(characterEntity, anim, keys)

}

func (m *Input) processMoving(char *donburi.Entry, anim *components.CurrentAnimationData, keys []ebiten.Key) {
	var shift gmath.Vec
	oldDirection := *direction.Direction.Get(char)
	newDirection := oldDirection

	if lo.Contains(keys, ebiten.KeyD) {
		shift.X++
		newDirection = direction.Right
	}
	if lo.Contains(keys, ebiten.KeyA) {
		shift.X--
		newDirection = direction.Left
	}
	if lo.Contains(keys, ebiten.KeyW) {
		shift.Y--
		newDirection = direction.Up
	}
	if lo.Contains(keys, ebiten.KeyS) {
		shift.Y++
		newDirection = direction.Down
	}

	if oldDirection != newDirection {
		direction.Direction.SetValue(char, newDirection)
	}

	if !shift.IsZero() {
		donburi.Add(char, components.Movement, &components.MovementData{Vec: shift})

		if anim.Entry != nil && anim.IsWalking && oldDirection != newDirection {
			anim.Entry.Remove()
			anim.Entry = nil
		}

		if anim.Entry == nil {
			anim.Entry = components.StartAnimation(char.World, *core.NewAnimationPlayer(
				time.Millisecond*600,
				core.WithOnFinish(func() { anim.Entry = nil }),
				core.WithAnimations(animations.NewSpritesheetAnimation(
					m.walkingAnimationMap,
					newDirection,
					time.Millisecond*600,
					components.Sprite.Get(char),
				)),
			))
			anim.IsWalking = true
		}
	} else {
		if anim.Entry != nil && anim.IsWalking {
			anim.Entry.Remove()
			anim.Entry = nil
		}

		components.Sprite.Get(char).Image = m.walkingAnimationMap.GetByDirection(newDirection, 0)
	}
}

func (m *Input) processNumbers(char *donburi.Entry, keys []ebiten.Key) {
	for i := range constants.DownPanelLength {
		if lo.Contains(keys, ebiten.Key(int(ebiten.Key1)+i)) {
			don.Create(char.World, gui.SelectCellRequest, &gui.SelectCellRequestData{CellNumber: i})
			return
		}
	}
}
