package main

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var assetsFolder embed.FS

func main() {
	fmt.Println("Break 'Em Free... made by Maxim Kozlov for the Ebitengine game jam 2026")

	ebiten.SetWindowSize(480, 270)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetFullscreen(true)

	player_spritesheet, _, imgerr := ebitenutil.NewImageFromFileSystem(assetsFolder, "assets/player.png")

	if imgerr != nil {
		log.Fatal(imgerr)
	}

	game := &Game{player: Player{ss: NewSpritesheet(player_spritesheet, 32, 32)}}

	gameErr := ebiten.RunGame(game)

	if gameErr != nil {
		if !errors.Is(gameErr, closeGame) {
			log.Fatal(gameErr)
		}
	}

}
