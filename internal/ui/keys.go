package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	Quit     key.Binding
	FlagCell key.Binding
	OpenCell key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.OpenCell, k.FlagCell, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {k.Up, k.Down, k.Left, k.Right},
        {k.OpenCell, k.FlagCell, k.Quit},
    }
}


var Keys = KeyMap{
	Up: key.NewBinding(
        key.WithKeys("up", "k", "w"),
        key.WithHelp("↑/k/w", "up"),
    ),
    Down: key.NewBinding(
        key.WithKeys("down", "j", "s"),
        key.WithHelp("↓/j/s", "down"),
    ),
	Left: key.NewBinding(
        key.WithKeys("left", "h", "a"),
        key.WithHelp("←/h/a", "left"),
    ),
	Right: key.NewBinding(
        key.WithKeys("right", "l", "d"),
        key.WithHelp("→/l/d", "right"),
    ),
	OpenCell: key.NewBinding(
        key.WithKeys(" ", "o"),
        key.WithHelp("space/o", "open cell"),
    ),
    FlagCell: key.NewBinding(
        key.WithKeys("f", "p"),
        key.WithHelp("f/p", "flag cell"),
    ),
    Quit: key.NewBinding(
        key.WithKeys("q", "ctrl+c"),
        key.WithHelp("q", "quit"),
    ),
}
