package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/Arjun259194/a-star/ds"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid      Grid
	pause     bool
	openSet   ds.IdxMap[Spot]
	closedSet ds.IdxMap[Spot]
	start     *Spot
	end       *Spot
	curr      *Spot
}

func (this *Game) Update() error {
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
		if spot.f < winnerSpot.f {
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

	neighboresIdxs := this.curr.getNeighbores()
	for _, idx := range neighboresIdxs {
		neighbore := this.grid.GetSpot(idx)

		exists := this.closedSet.Has(neighbore)
		if neighbore.wall || exists {
			continue
		}

		tempG := this.curr.g + 1

		isNewNode := !this.openSet.Has(neighbore)

		if isNewNode || tempG < neighbore.g {
			neighbore.g = tempG
			neighbore.h = neighbore.heuristic(*this.end)
			neighbore.f = neighbore.g + neighbore.h
			neighbore.prev = this.curr

			if isNewNode {
				this.openSet.Add(neighbore)
			}
		}

	}

	return nil
}

func (this Game) Draw(screen *ebiten.Image) {
	this.grid.Render(screen)

	this.end.draw(screen, color.RGBA{R: 255})

	for _, nidx := range this.end.getNeighbores() {
		n := this.grid.GetSpot(nidx)
		n.draw(screen, color.RGBA{R: 255, A: 0 })
	}

	temp := this.curr
	for temp != nil {
		temp.draw(screen, color.RGBA{B: 255})
		temp = temp.prev
	}

}

func (this *Game) Layout(outsideWidth, outSideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	grid := NewGrid(COLS, ROWS)

	i, j := rand.Intn(COLS-1), rand.Intn(ROWS-1)

	end := grid.GetSpotByIndex(i, j)

	game := &Game{
		grid:      grid,
		pause:     false,
		closedSet: ds.NewIdxMap[Spot](),
		openSet:   ds.NewIdxMap[Spot](),
		start:     grid.GetSpotByIndex(0, 0),
		end:       end,
		curr:      nil,
	}

	game.end.wall = false
	game.start.wall = false

	game.openSet.Add(game.start)

	ebiten.SetTPS(FRAMRATE)

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("A*")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
