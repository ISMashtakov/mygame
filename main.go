package main

import (
	"log"

	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func main() {
	resourceLoader := &resources.ResourceLoader{}

	world := donburi.NewWorld()
	gameObj := game.NewGame(
		world,
		[]core.ISystem{
			systems.NewInput(),
			systems.NewSwapSpriteByWalkingAnimation(60, resourceLoader),
			systems.NewMovement(),
		},
	)

	characterCreator := entities.NewCharacterCreator(world)

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
