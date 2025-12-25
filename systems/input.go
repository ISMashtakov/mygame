package systems

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/resources"
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
	hoeHittingAnimation *images.AnimationMap
}

func NewInput(resourcesLoader resources.IResourceLoader) *Input {
	return &Input{
		BaseSystem: core.BaseSystem{
			Codename:    InputCodename,
			NextSystems: []string{CollissionDetectorCodename, MovementCodename},
		},
		walkingAnimationMap: resourcesLoader.LoadAnimationMap(resources.AnimationCharacterWalking),
		hoeHittingAnimation: resourcesLoader.LoadAnimationMap(resources.AnimationCharacterHoeHitting),
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

	m.processNumbers(characterEntity, justPressedKeys)
	if m.processAction(characterEntity, anim, keys) {
		return
	}

	m.processMoving(characterEntity, anim, keys)
}

func (m *Input) processAction(char *donburi.Entry, anim *components.CurrentAnimationData, justPressedKeys []ebiten.Key) bool {
	if lo.Contains(justPressedKeys, ebiten.KeySpace) {
		panelEntity, ok := donburi.NewQuery(filter.Contains(gui.SelectedCell, gui.DownPanel)).First(char.World)
		if !ok {
			panic("panel not found")
		}

		selectedCell := gui.SelectedCell.Get(panelEntity)
		downPanel := gui.DownPanel.Get(panelEntity)

		item := downPanel.Items[selectedCell.CellNumber]
		if item != nil {
			switch item.GetType() {
			case items.Hoe:
				if anim.Entry != nil {
					components.DeleteAnimation(char.World, anim.Entry)
					anim.Entry = nil
				}

				anim.Entry = components.StartAnimation(char.World, *core.NewAnimationPlayer(
					time.Millisecond*600,
					core.WithOnFinish(func() {
						anim.Entry = nil
						point := components.Position.Get(char).Vec.Add(direction.GetDirectionVector(*direction.Direction.Get(char)).Mul(constants.TileSize))
						// сдвиг для красоты
						if *direction.Direction.Get(char) != direction.Down {
							point.Y += 10
						}

						don.CreateRequest(char.World, actions.GardenCreatingRequest, &actions.GardenCreatingRequestData{
							Point: point,
						})
					}),
					core.WithAnimations(animations.NewSpritesheetAnimation(
						m.hoeHittingAnimation,
						*direction.Direction.Get(char),
						time.Millisecond*600,
						components.Sprite.Get(char),
					)),
				))
				anim.IsWalking = false

			case items.Pickaxe:
				if anim.Entry != nil {
					components.DeleteAnimation(char.World, anim.Entry)
					anim.Entry = nil
				}

				anim.Entry = components.StartAnimation(char.World, *core.NewAnimationPlayer(
					time.Millisecond*600,
					core.WithOnFinish(func() {
						anim.Entry = nil
						point := components.Position.Get(char).Vec.Add(direction.GetDirectionVector(*direction.Direction.Get(char)).Mul(constants.TileSize))
						// сдвиг для красоты
						if *direction.Direction.Get(char) != direction.Down {
							point.Y += 10
						}
						don.CreateRequest(char.World, actions.PickaxeHitRequest, &actions.PickaxeHitRequestData{
							Point: point,
						})
					}),
					core.WithAnimations(animations.NewSpritesheetAnimation(
						m.hoeHittingAnimation,
						*direction.Direction.Get(char),
						time.Millisecond*600,
						components.Sprite.Get(char),
					)),
				))
				anim.IsWalking = false

			}
		}
	}

	return false
}

func (m *Input) processMoving(char *donburi.Entry, anim *components.CurrentAnimationData, keys []ebiten.Key) {
	var shift gmath.Vec
	oldDirection := *direction.Direction.Get(char)
	newDirection := oldDirection

	if lo.Contains(keys, ebiten.KeyD) {
		shift.X += 1
		newDirection = direction.Right
	}
	if lo.Contains(keys, ebiten.KeyA) {
		shift.X -= 1
		newDirection = direction.Left
	}
	if lo.Contains(keys, ebiten.KeyW) {
		shift.Y -= 1
		newDirection = direction.Up
	}
	if lo.Contains(keys, ebiten.KeyS) {
		shift.Y += 1
		newDirection = direction.Down
	}

	if oldDirection != newDirection {
		direction.Direction.SetValue(char, newDirection)
	}

	if !shift.IsZero() {
		donburi.Add(char, components.Movement, &components.MovementData{Vec: shift})

		if anim.Entry != nil && anim.IsWalking && oldDirection != newDirection {
			components.DeleteAnimation(char.World, anim.Entry)
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
			components.DeleteAnimation(char.World, anim.Entry)
			anim.Entry = nil
		}

		components.Sprite.Get(char).Image = m.walkingAnimationMap.GetByDirection(newDirection, 0)
	}
}

func (m *Input) processNumbers(char *donburi.Entry, keys []ebiten.Key) {
	for i := 0; i < constants.DownPanelLength; i++ {
		if lo.Contains(keys, ebiten.Key(int(ebiten.Key1)+i)) {
			don.CreateRequest(char.World, gui.SelectCellRequest, &gui.SelectCellRequestData{CellNumber: i})
			return
		}
	}
}
