package ui

import "github.com/charmbracelet/lipgloss"

var (
	colBlue   = lipgloss.Color(lipgloss.Color("#5C9EFF"))
	colGreen  = lipgloss.Color(lipgloss.Color("#5FD75F"))
	colRed    = lipgloss.Color(lipgloss.Color("#FF5F5F"))
	colPurple = lipgloss.Color(lipgloss.Color("#B45CFF"))
	colOrange = lipgloss.Color(lipgloss.Color("#FFAA00"))
	colCyan   = lipgloss.Color(lipgloss.Color("#00D7D7"))
	colPink   = lipgloss.Color(lipgloss.Color("#FF5FD7"))
	colGray   = lipgloss.Color(lipgloss.Color("#AAAAAA"))
	colWhite  = lipgloss.Color(lipgloss.Color("#ffffff"))
)

var numberStyles = map[int]lipgloss.Style{
	1: lipgloss.NewStyle().Foreground(colBlue),
	2: lipgloss.NewStyle().Foreground(colGreen),
	3: lipgloss.NewStyle().Foreground(colRed),
	4: lipgloss.NewStyle().Foreground(colPurple),
	5: lipgloss.NewStyle().Foreground(colOrange),
	6: lipgloss.NewStyle().Foreground(colCyan),
	7: lipgloss.NewStyle().Foreground(colPink),
	8: lipgloss.NewStyle().Foreground(colGray),
}

var (
	mineStyle = lipgloss.NewStyle().Foreground(colRed)
	flagStyle = lipgloss.NewStyle().Foreground(colWhite)
)

var boardBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder())
