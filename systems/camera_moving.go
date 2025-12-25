package systems

import (
	"log/slog"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	CameraMovingCodename = "CameraMovingCodename"
)

type CameraMoving struct {
	core.BaseSystem
}

func NewCameraMoving() *CameraMoving {
	return &CameraMoving{
		BaseSystem: core.BaseSystem{
			Codename:        CameraMovingCodename,
			PreviousSystems: []string{MovementCodename, CollissionDetectorCodename},
		},
	}
}

func (c *CameraMoving) Update(world donburi.World) {
	cameraEn, ok := donburi.NewQuery(filter.Contains(components.Position, components.Camera)).First(world)
	if !ok {
		slog.Debug("can not find camera")
	}

	cameraPos := components.Position.Get(cameraEn)

	characterEn, ok := donburi.NewQuery(filter.Contains(components.Position, components.Character)).First(world)
	if !ok {
		slog.Debug("can not find character")
	}

	characterPos := components.Position.Get(characterEn)

	cameraPos.Vec = characterPos.Vec
}
