package game

func (g *Game) MoveUp() {
	if g.curY > 0 {
		g.curY--
	}
}

func (g *Game) MoveDown() {
	if g.curY < g.Board.Height-1 {
		g.curY++
	}
}

func (g *Game) MoveLeft() {
	if g.curX > 0 {
		g.curX--
	}
}

func (g *Game) MoveRight() {
	if g.curX < g.Board.Width-1 {
		g.curX++
	}
}

func (g *Game) OpenCell() {

	if g.State == NotStarted {
		g.PlaceMines()
		g.State = StatePlaing
	}

	cell := g.SelectedCell()

	if cell.IsFlagged {
		return
	}

	if cell.NeighborMines != 0 {
		g.openNum(g.curX, g.curY)
	} else if !cell.IsMine && !cell.IsFlagged && !cell.IsRevealed {
		g.openAround(g.curX, g.curY, true)
	} else if cell.IsMine {
		g.GameOver()
	}
	// TODO fail game
}

func (g *Game) openAround(x, y int, first bool) {
	if g.Board.Grid[y][x].IsRevealed {
		return
	}
	g.Board.Grid[y][x].Open()

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := dx+x, dy+y
			if nx < 0 || nx >= g.Board.Width || ny < 0 || ny >= g.Board.Height {
				continue
			}
			if g.Board.Grid[ny][nx].NeighborMines == 0 && !g.Board.Grid[ny][nx].IsFlagged {
				g.openAround(nx, ny, false)
			}
			g.Board.Grid[ny][nx].Open()
		}
	}
}

func (g *Game) openNum(x, y int) {
	if !g.SelectedCell().IsRevealed {
		g.SelectedCell().IsRevealed = true
		return
	}

	flagsCount := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := dx+x, dy+y
			if nx < 0 || nx >= g.Board.Width || ny < 0 || ny >= g.Board.Height {
				continue
			}
			if g.Board.Grid[ny][nx].IsFlagged {
				flagsCount++
			}
		}
	}

	if flagsCount == g.SelectedCell().NeighborMines {
		g.openAround(g.curX, g.curY, false)
	}
}

// это для того что нелья было менять поля x y из вне (геттеры)
func (g *Game) CursorX() int { return g.curX }
func (g *Game) CursorY() int { return g.curY }
