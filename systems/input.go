package systems

import (
	"log"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
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
}

func NewInput() *Input {
	return &Input{
		core.BaseSystem{
			Codename:    InputCodename,
			NextSystems: []string{CollissionDetectorCodename, MovementCodename},
		},
	}
}

func (m *Input) Update(world donburi.World) error {
	characterEntity, ok := donburi.NewQuery(filter.Contains(components.Character)).First(world)
	if !ok {
		log.Println("can't found character")
		return nil
	}

	if characterEntity.HasComponent(actions.Action) {
		return nil
	}

	keys := inpututil.AppendPressedKeys(nil)
	m.processMoving(characterEntity, keys)

	justPressedKeys := inpututil.AppendJustPressedKeys(nil)

	if lo.Contains(justPressedKeys, ebiten.KeySpace) {
		donburi.Add(characterEntity, actions.Action, &actions.HoeHit)
	}

	cellEntity, ok := donburi.NewQuery(filter.Contains(gui.SelectedCell)).First(world)
	if !ok {
		log.Println("can't found interface cell entity")
		return nil
	}

	m.processNumbers(cellEntity, justPressedKeys)
	return nil
}

func (m *Input) processMoving(char *donburi.Entry, keys []ebiten.Key) {
	var shift gmath.Vec
	if lo.Contains(keys, ebiten.KeyD) {
		shift.X += 1
		direction.Direction.SetValue(char, direction.Right)
	}
	if lo.Contains(keys, ebiten.KeyA) {
		shift.X -= 1
		direction.Direction.SetValue(char, direction.Left)
	}
	if lo.Contains(keys, ebiten.KeyW) {
		shift.Y -= 1
		direction.Direction.SetValue(char, direction.Up)
	}
	if lo.Contains(keys, ebiten.KeyS) {
		shift.Y += 1
		direction.Direction.SetValue(char, direction.Down)
	}

	if !shift.IsZero() {
		donburi.Add(char, components.MovementRequest, &components.MovementRequestData{Vec: shift})
	}
}

func (m *Input) processNumbers(en *donburi.Entry, keys []ebiten.Key) {
	for i := 0; i < 9; i++ {
		if lo.Contains(keys, ebiten.Key(int(ebiten.Key1)+i)) {
			donburi.Add(en, gui.SelectCellRequest, &gui.SelectCellRequestData{CellNumber: i})
			return
		}
	}
}
