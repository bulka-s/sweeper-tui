package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	Quit     key.Binding
	FlagCell key.Binding
	OpenCell key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
        key.WithKeys("up", "k", "w"),
        key.WithHelp("↑/k", "up"),
    ),
    Down: key.NewBinding(
        key.WithKeys("down", "j", "s"),
        key.WithHelp("↓/j", "down"),
    ),
	Left: key.NewBinding(
        key.WithKeys("left", "h", "a"),
        key.WithHelp("←/h", "left"),
    ),
	Right: key.NewBinding(
        key.WithKeys("right", "l", "d"),
        key.WithHelp("→/l", "right"),
    ),
	OpenCell: key.NewBinding(
        key.WithKeys(" ", "o"),
        key.WithHelp("space", "open cell"),
    ),
    FlagCell: key.NewBinding(
        key.WithKeys("f", "p"),
        key.WithHelp("F", "flag cell"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("q", "ctrl+c"),
        key.WithHelp("q", "quit"),
    ),
}
