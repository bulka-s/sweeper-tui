package main

import (
	"fmt"
	"minesweeper/internal/game"
	"minesweeper/internal/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	g := game.NewGame(40, 20, 100)

	p := tea.NewProgram(ui.NewModel(g))
	p.Run()
}
