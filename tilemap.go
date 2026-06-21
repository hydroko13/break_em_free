package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileChunk struct {
	chunk_data  *[8 * 8]uint
	x           int
	y           int
	tile_size   int
	chunk_image *ebiten.Image
}

func NewEmptyTileChunk(x, y, tile_size int) TileChunk {
	return TileChunk{
		chunk_data:  &[8 * 8]uint{},
		x:           x,
		y:           y,
		chunk_image: ebiten.NewImage(8*tile_size, 8*tile_size),
		tile_size:   tile_size,
	}
}

func (c *TileChunk) SetTileAt(x, y int, id uint) {
	c.chunk_data[y*8+x] = id
}

func (c *TileChunk) Redraw(ss Spritesheet) {
	for tile_idx, tile_id := range c.chunk_data {
		if tile_id != 0 {

			grid_x, grid_y := tile_idx%8, tile_idx/8

			tile_x, tile_y := grid_x*c.tile_size, grid_y*c.tile_size
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile_x), float64(tile_y))
			c.chunk_image.DrawImage(ss.GetFrame(0, int(tile_id-1)), &op)
		}

	}
}

type Tilemap struct {
	tile_ss     Spritesheet // Spritesheet for tileset for this tilemap
	tile_size   int
	tile_chunks []TileChunk
}

func NewTilemap(tile_ss Spritesheet) Tilemap {
	if tile_ss.frame_width != tile_ss.frame_height {
		// assert that the frame_width must match frame_height
		panic(
			fmt.Sprintf(`Tiles must be square, the frame width should exactly match the frame height
In this case the frame size was (%d, %d)`,
				tile_ss.frame_width, tile_ss.frame_height,
			))

	}
	return Tilemap{
		tile_ss:     tile_ss,
		tile_size:   tile_ss.frame_width,
		tile_chunks: make([]TileChunk, 0),
	}
}

func (tilemap Tilemap) Draw(screen *ebiten.Image, camera Camera) {
	for _, tile_chunk := range tilemap.tile_chunks {

		wx, wy := (tile_chunk.x * tilemap.tile_size * 8), (tile_chunk.y * tilemap.tile_size * 8)
		op := ebiten.DrawImageOptions{}
		x, y := camera.OffsetPoint(float64(wx), float64(wy))
		op.GeoM.Translate(x, y)
		screen.DrawImage(tile_chunk.chunk_image, &op)

	}

}
