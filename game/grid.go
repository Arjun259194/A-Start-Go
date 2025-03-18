package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grid [][]Spot

func NewGrid(cols, rows int, wallRate float32, size int) Grid {
	grid := make(Grid, cols)
	for i := range grid {
		grid[i] = make([]Spot, rows)
		for j := range grid[i] {
			grid[i][j] = NewSpot(i, j, rand.Float32() < wallRate, float32(size))
		}
	}
	return grid
}

func (this Grid) GetSpotByIndex(i, j int) *Spot {
	return &this[i][j]
}

func (this Grid) GetSpot(index SpotIndex) *Spot {
	i, j := index.Get()
	return &this[i][j]
}

func (this Grid) Render(screen *ebiten.Image) {
	for _, row := range this {
		for _, spot := range row {
			if spot.Wall {
				spot.Draw(screen, color.White)
			}
		}
	}
}
