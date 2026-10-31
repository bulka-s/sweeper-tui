package game

import (
	"math/rand/v2"
)

type Game struct {
	Board      Board
	MinesCount int
	curX       int
	curY       int
	State      GameState
}

type GameState int

const (
	StatePlaying GameState = iota
	StateWon
	StateGameOver
	NotStarted
)

type point struct{ X, Y int }

func NewGame(width, height, minesCount int) *Game {
	grid := make([][]Cell, height)
	for i := range grid {
		grid[i] = make([]Cell, width)
	}

	return &Game{
		Board: Board{
			Width:  width,
			Height: height,
			Grid:   grid,
		},
		MinesCount: minesCount,
		curX:       width / 2,
		curY:       height / 2,
		State:      NotStarted,
	}
}

func (g *Game) PlaceMines(x, y int) {
	positions := []point{}

	for i := 0; i < g.Board.Height; i++ {
		for j := 0; j < g.Board.Width; j++ {
			if absInt(j-x) <= 1 && absInt(i-y) <= 1 {
				continue
			}
			positions = append(positions, point{j, i})
		}
	}

	if len(positions) < g.MinesCount {
		return
	}

	rand.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	for i := 0; i < g.MinesCount; i++ {
		g.Board.Grid[positions[i].Y][positions[i].X].IsMine = true
	}

	g.Board.setNeighborsCount()
}

func (g *Game) SelectedCell() *Cell {
	return &g.Board.Grid[g.curY][g.curX]
}

func (g *Game) GameOver() {
	g.Board.OpenAllMines()
	g.State = StateGameOver
}
