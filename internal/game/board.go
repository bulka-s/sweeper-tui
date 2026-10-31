package game

type Cell struct {
	IsMine        bool
	IsRevealed    bool
	IsFlagged     bool
	NeighborMines int
}

type Board struct {
	Width  int
	Height int

	Grid [][]Cell
}

func (b Board) OpenAllMines() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Grid[y][x].IsMine {
				b.Grid[y][x].Open()
			}
		}
	}
}

func (c *Cell) Open() {
	if !c.IsRevealed {
		c.IsRevealed = true
	}
}

func (c *Cell) ToggleFlag() {
	if !c.IsRevealed {
		c.IsFlagged = !c.IsFlagged
	}
}

func (b *Board) setNeighborsCount() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Grid[y][x].IsMine {
				continue
			}
			b.Grid[y][x].NeighborMines = countNeighbors(x, y, b)
		}
	}
}

func countNeighbors(x, y int, b *Board) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := dx+x, dy+y
			if nx < 0 || nx >= b.Width || ny < 0 || ny >= b.Height {
				continue
			}
			if b.Grid[ny][nx].IsMine {
				count++
			}
		}
	}
	return count
}
