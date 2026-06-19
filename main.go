package main

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assetsFolder embed.FS



func main() {
	fmt.Println("Break 'Em Free... made by Maxim Kozlov for the Ebitengine game jam 2026")

	ebiten.SetWindowSize(480, 270)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetFullscreen(true)

	game := &Game{}

	gameErr := ebiten.RunGame(game)

	if gameErr != nil {
		if !errors.Is(gameErr, closeGame) {
			log.Fatal(gameErr)
		}	
	}





}