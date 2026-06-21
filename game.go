package main

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var closeGame error = errors.New("Game closed")

type Game struct {
	cam     Camera
	player  Player
	tilemap Tilemap
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{59, 58, 58, 255})

	g.player.Draw(screen, g.cam)
	g.tilemap.Draw(screen, g.cam)

}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return closeGame
	}

	g.player.Update()

	g.cam.x += (g.player.x - g.cam.x) * DELTA * 4.2
	g.cam.y += (g.player.y - g.cam.y) * DELTA * 4.2

	return nil
}

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return WIDTH, HEIGHT
}
