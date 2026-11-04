package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) renderBoard() string {
	var b strings.Builder
	for y, row := range m.game.Board.Grid {
		for x, c := range row {
			var style lipgloss.Style
			var text string

			switch {
			case c.IsMine && c.IsRevealed:
				style = mineStyle
				text = "@"
			case c.IsFlagged && !c.IsRevealed:
				style = flagStyle
				text = "⚑"
			case c.IsRevealed && c.NeighborMines > 0:
				style = numberStyles[c.NeighborMines]
				text = strconv.Itoa(c.NeighborMines)
			case c.IsRevealed:
				style = lipgloss.NewStyle()
				text = " "
			default:
				style = lipgloss.NewStyle()
				text = "."
			}

			selected := x == m.game.CursorX() && y == m.game.CursorY()
			if selected {
				style = style.Reverse(true)
			}

			b.WriteString(style.Render(text))
			b.WriteRune(' ')
			m.game.CursorX()
		}
		if y < m.game.Board.Height-1 {
			b.WriteRune('\n')
		}
	}
	return boardBorderStyle.Render(b.String())
}

func (m model) renderTopBar() string {
	minesLeft := m.game.MinesCount - m.game.FlagsPlaced
	topBar := fmt.Sprintf("Mines: %d/%d | Pos: [%d, %d]", minesLeft, m.game.MinesCount, m.game.CursorX(), m.game.CursorY())
	return topBar
}
