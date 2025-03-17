package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Spot struct {
	index   SpotIndex
	f, g, h float32
	wall    bool
	prev    *Spot
}

func newSpot(i, j int) Spot {
	return Spot{
		index: NewSpotIndex(i, j),
		f:     0,
		g:     0,
		h:     0,
		wall:  rand.Float32() < WALL_RATE,
		prev:  nil,
	}
}

func (this Spot) draw(screen *ebiten.Image, clr color.Color) {
	i, j := this.index.Get()
	x, y := float32(i*SIZE), float32(j*SIZE)
	size := float32(SIZE)
	vector.DrawFilledRect(screen, x, y, size, size, clr, false)
}

func (this Spot) heuristic(spot Spot) float32 {
	x1, y1 := this.index.Get()
	x2, y2 := spot.index.Get()

	x := math.Pow(float64(x1-x2), 2)
	y := math.Pow(float64(y1-y2), 2)

	dist := math.Sqrt(x + y)

	return float32(dist)
}

func (this Spot) getNeighbores() []SpotIndex {
	i, j := this.index.Get()

	neighbores := make([]SpotIndex, 8)

	if i < COLS-1 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j))
	}
	if i > 0 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j))
	}
	if j < ROWS-1 {
		neighbores = append(neighbores, NewSpotIndex(i, j+1))
	}
	if j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i, j-1))
	}

	//diagnal
	if i > 0 && j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j-1))
	}
	if i < COLS-1 && j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j-1))
	}
	if i > 0 && j < ROWS-1 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j+1))
	}
	if i < COLS-1 && j < ROWS-1 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j+1))
	}

	return neighbores
}
