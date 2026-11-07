package ui

import (
	"fmt"
	"minesweeper/internal/game"
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
				text = "X"
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
	return m.renderBorder(b.String())
}

func (m model) renderBorder(content string) string {
	var borderStyle lipgloss.Style 

	switch m.game.State{
	case game.StateGameOver:
		borderStyle = boardBorderStyle.BorderForeground(lipgloss.Color("#b50404"))
	case game.StateWon:
		borderStyle = boardBorderStyle.BorderForeground(lipgloss.Color("#04b527"))
	default:
		borderStyle = boardBorderStyle.BorderForeground(lipgloss.Color("#b7b7b7"))
	}

	return borderStyle.Render(content)
}

func (m model) renderTopBar() string {
	minesLeft := m.game.MinesCount - m.game.FlagsPlaced
	topBar := fmt.Sprintf("Mines: %d/%d | Pos: [%d, %d] | TimeL: %d", minesLeft, m.game.MinesCount, m.game.CursorX(), m.game.CursorY(), m.seconds)
	return topBar
}
