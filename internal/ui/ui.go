package ui

import (
	"minesweeper/internal/game"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	game   *game.Game
	width  int
	height int
	help   help.Model
}

func NewModel(g *game.Game) *model {
	return &model{
		game: g,
		help: help.New(),
	}
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
	case msg.String() == "?":
		m.help.ShowAll = !m.help.ShowAll

	case key.Matches(msg, Keys.Quit):
		return m, tea.Quit
	case key.Matches(msg, Keys.Up):
		m.game.MoveUp()
	case key.Matches(msg, Keys.Down):
		m.game.MoveDown()
	case key.Matches(msg, Keys.Left):
		m.game.MoveLeft()
	case key.Matches(msg, Keys.Right):
		m.game.MoveRight()
	case key.Matches(msg, Keys.OpenCell):
		m.game.OpenCell()
	case key.Matches(msg, Keys.FlagCell):
		if m.game.State == game.StatePlaying {
			m.game.FlagSetter()
		}
	case key.Matches(msg, Keys.RestartGame):
		m.game = game.NewGame(m.game.Board.Width, m.game.Board.Height, m.game.MinesCount)
	}

	return m, nil
}

func (m *model) View() string {

	topBar := m.renderTopBar()
	board := m.renderBoard()
	helpStr := m.help.View(Keys)

	content := lipgloss.JoinVertical(lipgloss.Center, topBar, board, helpStr)

	if m.width == 0 || m.height == 0 {
		return content
	}

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}
