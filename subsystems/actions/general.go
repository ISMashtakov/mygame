package actions

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/items"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

func startHitAnimation(hoeHittingAnimation *images.AnimationMap, characterEntity *donburi.Entry, onFinish func(point gmath.Vec)) {
	anim := components.CurrentAnimation.Get(characterEntity)

	if anim.Entry != nil {
		anim.Entry.Remove()
		anim.Entry = nil
	}

	anim.Entry = components.StartAnimation(characterEntity.World, *core.NewAnimationPlayer(
		time.Millisecond*600,
		core.WithOnFinish(func() {
			anim.Entry = nil
			point := components.Position.Get(characterEntity).Vec.Add(direction.GetDirectionVector(*direction.Direction.Get(characterEntity)).Mul(constants.TileSize))
			// сдвиг для красоты
			if *direction.Direction.Get(characterEntity) != direction.Down {
				point.Y += 10
			}

			onFinish(point)
		}),
		core.WithAnimations(animations.NewSpritesheetAnimation(
			hoeHittingAnimation,
			*direction.Direction.Get(characterEntity),
			time.Millisecond*600,
			components.Sprite.Get(characterEntity),
		)),
	))
	anim.IsWalking = false
}

func spaceIsPressed() bool {
	justPressedKeys := inpututil.AppendJustPressedKeys(nil)
	return lo.Contains(justPressedKeys, ebiten.KeySpace)
}

func getCursorPos(world donburi.World) gmath.Vec {
	cursorX, cursorY := ebiten.CursorPosition()
	cursorPos := gmath.Vec{X: float64(cursorX), Y: float64(cursorY)}

	camera, ok := donburi.NewQuery(filter.Contains(components.Position, components.Camera)).First(world)
	if !ok {
		panic("not found camera")
	}

	cameraPos := components.Position.Get(camera)

	return cursorPos.Add(cameraPos.Vec).Sub(constants.TargetLayout.Mulf(0.5))
}

func getSelectedItem(world donburi.World) items.IItem {
	panelEntity, ok := donburi.NewQuery(filter.Contains(gui.SelectedCell, gui.DownPanel)).First(world)
	if !ok {
		panic("panel not found")
	}

	selectedCell := gui.SelectedCell.Get(panelEntity)
	downPanel := gui.DownPanel.Get(panelEntity)

	return downPanel.GetItem(selectedCell.CellNumber)
}
