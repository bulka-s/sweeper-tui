package ui

import "github.com/charmbracelet/lipgloss"

var numberStyles = map[int]lipgloss.Style{
	1: lipgloss.NewStyle().Foreground(lipgloss.Color("#5C9EFF")), // синий
	2: lipgloss.NewStyle().Foreground(lipgloss.Color("#5FD75F")), // зелёный
	3: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F5F")), // красный
	4: lipgloss.NewStyle().Foreground(lipgloss.Color("#B45CFF")), // фиолетовый
	5: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFAA00")), // оранжевый
	6: lipgloss.NewStyle().Foreground(lipgloss.Color("#00D7D7")), // бирюзовый
	7: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5FD7")), // розовый
	8: lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA")), // серый
}

var (
	mineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
	flagStyle = lipgloss.NewStyle().Background(lipgloss.Color("#d2e500"))
)

var boardBorderStyle = lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#b7b7b7")) 
