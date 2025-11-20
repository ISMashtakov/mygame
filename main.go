package main

import (
	"log"

	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func main() {
	resourceLoader := resources.NewResourceLoader()
	resourceLoader.Preload()

	renderer := *game.NewRenderer()
	// renderer.DrawColliders = true

	world := donburi.NewWorld()
	gameObj := game.NewGame(
		renderer,
		world,
		[]core.ISystem{
			systems.NewInput(),
			systems.NewSwapSpriteByWalkingAnimation(60, resourceLoader),
			systems.NewCollisionDetector(),
			systems.NewMovement(),
		},
	)

	// Entity creatores
	characterCreator := entities.NewCharacterCreator()
	grassCreator := background.NewGrassCreator(resourceLoader)
	stoneCreator := entities.NewStoneCreator(resourceLoader)

	// ----------
	worldBuilder := game.NewWorldBuilder(*grassCreator, *stoneCreator)

	err := worldBuilder.Build(world)
	if err != nil {
		log.Fatalf("can't create character: %s", err)
	}

	_, err = characterCreator.Create(world)
	if err != nil {
		log.Fatalf("can't create character: %s", err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(gameObj); err != nil {
		log.Fatal(err)
	}
}
