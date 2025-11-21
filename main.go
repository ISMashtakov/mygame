package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	builder := Builder{}

	builder.Resources()
	builder.Renderer()
	builder.Systems()
	builder.World()
	builder.Entities()

	builder.RunGame()
}
