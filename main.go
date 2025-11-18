package main

import (
	"log"

	"github.com/ISMashtakov/mygame/character"
	"github.com/ISMashtakov/mygame/character/input"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/physics"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

func main() {
	world := ecs.NewWorld()
	gameObj := game.NewGame(
		&world,
		[]core.ISystem{
			input.NewInput(&world),
			physics.NewMovement(&world),
		},
	)

	resourceLoader := &resources.ResourceLoader{}

	characterCreator := character.NewCharacterCreator(&world, resourceLoader)

	_, err := characterCreator.Create()
	if err != nil {
		log.Fatalf("can't create character: %s", err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(gameObj); err != nil {
		log.Fatal(err)
	}
}
