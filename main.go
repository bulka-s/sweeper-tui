package main

import (
	"minesweeper/internal/game"
	"minesweeper/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	g := game.NewGame(40, 20, 100)

	p := tea.NewProgram(ui.NewModel(g))
	p.Run()
}
