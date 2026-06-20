package main

type Camera struct {
	x float64
	y float64
}

func (cam Camera) OffsetPoint(x, y float64) (float64, float64) {
	return x - cam.x + (float64(WIDTH) / 2), y - cam.y + (float64(HEIGHT) / 2)
}
