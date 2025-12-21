package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1920, 1080)

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
