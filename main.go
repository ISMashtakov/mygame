package main

import (
	"log"

	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func main() {
	gameObj := game.Game{
		World: donburi.NewWorld(),
	}

	resourceLoader := &resources.ResourceLoader{}

	_, err := entities.CreateCharacter(gameObj.World, resourceLoader)
	if err != nil {
		log.Fatalf("can't create character: %s", err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&gameObj); err != nil {
		log.Fatal(err)
	}
}
