package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)

	builder := Builder{}

	builder.Resources()
	builder.Renderer()
	builder.World()
	builder.GUI()
	builder.Entities()
	builder.Systems()

	// builder.Debug()

	builder.RunGame()
}
