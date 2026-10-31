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

	if g.State == StateGameOver {
		return
	}

	if g.State == NotStarted {
		g.PlaceMines(g.curX, g.curY)
		g.State = StatePlaying
	}

	cell := g.SelectedCell()

	if cell.IsFlagged {
		return
	}

	if cell.IsMine {
		g.GameOver()
		return
	}

	if cell.IsRevealed {
		g.openNum(g.curX, g.curY)
		return
	}

	if cell.NeighborMines == 0 {
		g.openAround(g.curX, g.curY)
	} else {
		g.revealOne(g.curX, g.curY)
	}
}

func (g *Game) revealOne(x, y int) {
	if x < 0 || x >= g.Board.Width || y < 0 || y >= g.Board.Height {
		return
	}
	cell := &g.Board.Grid[y][x]
	if cell.IsRevealed || cell.IsFlagged {
		return
	}
	if cell.IsMine {
		g.GameOver()
		return
	}
	cell.IsRevealed = true
}

func (g *Game) openAround(x, y int) {
	if x < 0 || x >= g.Board.Width || y < 0 || y >= g.Board.Height {
		return
	}

	cell := &g.Board.Grid[y][x]
	if cell.IsRevealed || cell.IsFlagged {
		return
	}

	if cell.IsMine {
		g.GameOver()
		return
	}

	cell.IsRevealed = true

	if cell.NeighborMines != 0 {
		return
	}

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			g.openAround(x+dx, y+dy)
		}
	}
}

func (g *Game) openNum(x, y int) {
	if x < 0 || x >= g.Board.Width || y < 0 || y >= g.Board.Height {
		return
	}
	cell := &g.Board.Grid[y][x]
	if !cell.IsRevealed || cell.NeighborMines == 0 {
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

	if flagsCount != cell.NeighborMines {
		return
	}

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			g.openAround(x+dx, y+dy)
		}
	}

}

// это для того что нелья было менять поля x y из вне (геттеры)
func (g *Game) CursorX() int { return g.curX }
func (g *Game) CursorY() int { return g.curY }
