package main

import "sync"
import "time"
import "github.com/aquilax/go-perlin"
import "slices"

type WorldGenerator struct {
	noise1 *perlin.Perlin
}

type ChunkPos struct {
	x int
	y int
}

func (wg *WorldGenerator) StartGenerating(tilemap *Tilemap, tilemap_mutex *sync.Mutex) {
	generated_positions := make([]ChunkPos, 0, 15)

	for {
		for x := -4; x < 4; x += 1 {
			for y := -4; y < 4; y += 1 {
				if !slices.Contains(generated_positions, ChunkPos{x: x, y: y}) {
					generated_positions = append(generated_positions, ChunkPos{x: x, y: y})

					chunk := NewEmptyTileChunk(x, y, 16)

					for cx := 0; cx < 8; cx += 1 {
						h := ((wg.noise1.Noise1D(float64(cx+(x*8))/12)+1)/2.0*6.5 + 2)
						for cy := 0; cy < 8; cy += 1 {
							if float64(cy+(y*8)) > h {
								if float64(cy+(y*8)) - 1 < h {
									chunk.SetTileAt(cx, cy, 1)
								} else {
									chunk.SetTileAt(cx, cy, 2)
								}
								
							}

						}
					}

					tilemap_mutex.Lock()
					chunk.Redraw(tilemap.tile_ss)
					tilemap.tile_chunks = append(tilemap.tile_chunks, chunk)
					tilemap_mutex.Unlock()
					time.Sleep(time.Millisecond * 8)

				}

			}
		}
	}

}
