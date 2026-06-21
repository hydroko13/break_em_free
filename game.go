package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"sync"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var closeGame error = errors.New("Game closed")

type Game struct {
	cam           Camera
	player        Player
	tilemap       Tilemap
	tilemap_mutex sync.Mutex
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{59, 58, 58, 255})

	g.player.Draw(screen, g.cam)

	g.tilemap_mutex.Lock()

	g.tilemap.Draw(screen, g.cam)

	g.tilemap_mutex.Unlock()

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
