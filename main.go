package main

import (
	"github.com/ISMashtakov/mygame/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(constants.TargetLayout.ToStd().X*2, constants.TargetLayout.ToStd().Y*2)

	builder := Builder{}

	builder.Resources()
	builder.Renderer()
	builder.ItemsFactory()
	builder.GUI()
	builder.World()
	builder.Entities()
	builder.Systems()

	// builder.Debug()

	builder.RunGame()
}
