package ui

import (
	"minesweeper/internal/game"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	game   *game.Game
	width  int
	height int
}

func NewModel(g *game.Game) *model {
	return &model{game: g}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		return m.handleKey(msg)
	}
	return m, nil
}

func (m *model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Quit):
		return m, tea.Quit
	case key.Matches(msg, keys.Up):
		m.game.MoveUp()
	case key.Matches(msg, keys.Down):
		m.game.MoveDown()
	case key.Matches(msg, keys.Left):
		m.game.MoveLeft()
	case key.Matches(msg, keys.Right):
		m.game.MoveRight()
	case key.Matches(msg, keys.OpenCell):
		m.game.OpenCell()
	case key.Matches(msg, keys.FlagCell):
		m.game.SelectedCell().ToggleFlag()
	}

	return m, nil
}

func (m *model) View() string {
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
				text = " "
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
	boardWithBorder := boardBorderStyle.Render(b.String())

	if m.width == 0 || m.height == 0 {
		return boardWithBorder
	}

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		boardWithBorder,
	)
}
