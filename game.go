package main

import (
	"image/color"
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)


var closeGame error = errors.New("Game closed")

type Game struct {
	
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}

func (g *Game) Update() error {


	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return closeGame
	}

	return nil
}

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return 480, 270
}
