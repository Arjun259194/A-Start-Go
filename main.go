package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/Arjun259194/a-star/ds"
	"github.com/Arjun259194/a-star/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid      game.Grid
	pause     bool
	openSet   ds.IdxMap[*game.Spot]
	closedSet ds.IdxMap[*game.Spot]
	start     *game.Spot
	end       *game.Spot
	curr      *game.Spot
}

func (this *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		this.pause = false
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if this.pause {
		return nil
	}

	if this.openSet.Len() <= 0 {
		fmt.Println("No solution found.....")
		this.pause = true
		return nil
	}

	winnerIdx := 0
	winnerSpot := this.openSet.Get(winnerIdx)
	for i, spot := range this.openSet.Iter() {
		if spot.F < winnerSpot.F {
			winnerIdx = i
			winnerSpot = spot
		}
	}

	this.curr = this.openSet.Get(winnerIdx)
	if this.curr == this.end {
		this.pause = true
		fmt.Println("DONE!!")
		return nil
	}

	this.openSet.Remove(this.curr)
	this.closedSet.Add(this.curr)

	neighboresIdxs := this.curr.GetNeighbores(COLS, ROWS)
	for _, idx := range neighboresIdxs {
		neighbore := this.grid.GetSpot(idx)

		exists := this.closedSet.Has(neighbore)
		if neighbore.Wall || exists {
			continue
		}

		tempG := this.curr.G + 1

		isNewNode := !this.openSet.Has(neighbore)

		if isNewNode || tempG < neighbore.G {
			neighbore.G = tempG
			neighbore.H = neighbore.Heuristic(*this.end)
			neighbore.F = neighbore.G + neighbore.H
			neighbore.Prev = this.curr

			if isNewNode {
				this.openSet.Add(neighbore)
			}
		}

	}

	return nil
}

func (this Game) Draw(screen *ebiten.Image) {
	this.grid.Render(screen)

	this.end.Draw(screen, color.RGBA{R: 255})

	temp := this.curr
	for temp != nil {
		temp.DrawPath(screen, color.RGBA{B: 255, G: 255})
		temp = temp.Prev
	}

}

func (this *Game) Layout(outsideWidth, outSideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	grid := game.NewGrid(COLS, ROWS, WALL_RATE, SIZE)

	i, j := rand.Intn(COLS-1), rand.Intn(ROWS-1)

	end := grid.GetSpotByIndex(i, j)

	game := &Game{
		grid:      grid,
		pause:     true,
		closedSet: ds.NewIdxMap[*game.Spot](),
		openSet:   ds.NewIdxMap[*game.Spot](),
		start:     grid.GetSpotByIndex(0, 0),
		end:       end,
		curr:      nil,
	}

	game.end.Wall = false
	game.start.Wall = false

	game.openSet.Add(game.start)

	ebiten.SetTPS(FRAMRATE)

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("A*")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
