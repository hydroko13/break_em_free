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

	var err error

	var player_spritesheet *ebiten.Image

	player_spritesheet, _, err = ebitenutil.NewImageFromFileSystem(assetsFolder, "assets/player.png")

	if err != nil {
		log.Fatal(err)
	}

	var groundtileset_spritesheet *ebiten.Image

	groundtileset_spritesheet, _, err = ebitenutil.NewImageFromFileSystem(assetsFolder, "assets/groundtileset.png")
	

	
	if err != nil {
		log.Fatal(err)
	}


	tilemap := NewTilemap(NewSpritesheet(groundtileset_spritesheet, 16, 16))

	tile_chunk1 := NewEmptyTileChunk(0, 1)
	tile_chunk2 := NewEmptyTileChunk(1, 1)
	tile_chunk1.chunk_data[0] = 1
	tile_chunk2.chunk_data[0] = 1

	tilemap.tile_chunks = append(tilemap.tile_chunks, tile_chunk1, tile_chunk2)

	game := &Game{player: Player{ss: NewSpritesheet(player_spritesheet, 32, 32)}, tilemap: tilemap}

	gameErr := ebiten.RunGame(game)

	if gameErr != nil {
		if !errors.Is(gameErr, closeGame) {
			log.Fatal(gameErr)
		}
	}

}
