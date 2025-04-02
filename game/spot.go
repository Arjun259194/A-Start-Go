package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SpotIndex struct{ i, j int }

func NewSpotIndex(i, j int) SpotIndex { return SpotIndex{i, j} }

func (this SpotIndex) Get() (int, int) {
	return this.i, this.j
}

func (this SpotIndex) GetFloat32() (float32, float32) {
	return float32(this.i), float32(this.j)
}

type Spot struct {
	index   SpotIndex
	F, G, H float32
	Wall    bool
	Prev    *Spot
	size    float32
}

func NewSpot(i, j int, isWall bool, size float32) Spot {
	return Spot{
		index: NewSpotIndex(i, j),
		F:     0,
		G:     0,
		H:     0,
		Wall:  isWall,
		Prev:  nil,
		size:  size,
	}
}

func (this Spot) Draw(screen *ebiten.Image, clr color.Color) {
	i, j := this.index.GetFloat32()
	size := this.size
	x, y := i*size, j*size
	vector.DrawFilledCircle(screen, x+size/2, y+size/2, size/3, clr, true)
}

func (this Spot) DrawPath(screen *ebiten.Image, clr color.Color) {
	if this.Prev == nil {
		return
	}
	size := this.size
	i0, j0 := this.index.GetFloat32()
	x0, y0 := i0*size+size/2, j0*size+size/2

	i1, j1 := this.Prev.index.GetFloat32()
	x1, y1 := i1*size+size/2, j1*size+size/2

	strokeWidth := size/2 - 1
	vector.StrokeLine(screen, x0, y0, x1, y1, strokeWidth, clr, false)
}


// MIGHT IMPLEMENT IN FUTURE
// function heuristic(node) =
//     dx = abs(node.x - goal.x)
//     dy = abs(node.y - goal.y)
//     return D * (dx + dy) + (D2 - 2 * D) * min(dx, dy)

func (this Spot) Heuristic(spot Spot) float32 {
	x1, y1 := this.index.GetFloat32()
	x2, y2 := spot.index.GetFloat32()

	x := math.Pow(float64(x1-x2), 2)
	y := math.Pow(float64(y1-y2), 2)

	dist := math.Sqrt(x + y)

	return float32(dist)
}

func (this Spot) GetNeighbores(cols, rows int) []SpotIndex {
	i, j := this.index.Get()

	neighbores := make([]SpotIndex, 8)

	if i < cols-1 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j))
	}
	if i > 0 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j))
	}
	if j < rows-1 {
		neighbores = append(neighbores, NewSpotIndex(i, j+1))
	}
	if j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i, j-1))
	}

	//diagnal
	if i > 0 && j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j-1))
	}
	if i < cols-1 && j > 0 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j-1))
	}
	if i > 0 && j < rows-1 {
		neighbores = append(neighbores, NewSpotIndex(i-1, j+1))
	}
	if i < cols-1 && j < rows-1 {
		neighbores = append(neighbores, NewSpotIndex(i+1, j+1))
	}

	return neighbores
}
