package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpotIndex struct{ i, j int }

func NewSpotIndex(i, j int) SpotIndex { return SpotIndex{i, j} }

func (this SpotIndex) Get() (int, int) {
	return this.i, this.j
}

type Grid [][]Spot

func NewGrid(cols, rows int) Grid {
	grid := make(Grid, cols)
	for i := range grid {
		grid[i] = make([]Spot, rows)
		for j := range grid[i] {
			grid[i][j] = newSpot(i, j)
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
			if spot.wall {
				spot.draw(screen, color.White)
			} 
		}
	}
}
