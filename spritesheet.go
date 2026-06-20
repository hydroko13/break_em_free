package main

import "github.com/hajimehoshi/ebiten/v2"

type Spritesheet struct {
	spritesheet_image *ebiten.Image
	frame_width       int
	frame_height      int
}

func NewSpritesheet(img *ebiten.Image, width, height int) Spritesheet {
	return Spritesheet{
		spritesheet_image: img,
		frame_width:       width,
		frame_height:      height,
	}
}

func (ss Spritesheet) GetFrame(row int, col int) *ebiten.Image {
	frame := ebiten.NewImage(ss.frame_width, ss.frame_height)
	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(-col*ss.frame_width), float64(-row*ss.frame_height))
	frame.DrawImage(ss.spritesheet_image, &ops)
	return frame
}
