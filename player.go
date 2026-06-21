package main

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	ss      Spritesheet
	x, y    float64
	vely    float64
	flipped bool
}

func (p Player) Draw(screen *ebiten.Image, cam Camera) {
	op := ebiten.DrawImageOptions{}
	x, y := cam.OffsetPoint(p.x, p.y)

	if p.flipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(32, 0)
	} else {
		op.GeoM.Scale(1, 1)
	}

	op.GeoM.Translate(x, y)

	screen.DrawImage(p.ss.GetFrame(0, 0), &op)
}

func (p *Player) Update() {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= DELTA * 95
		p.flipped = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += DELTA * 95
		p.flipped = true
	}

	p.y += p.vely * DELTA
	p.vely += DELTA * 50

}
