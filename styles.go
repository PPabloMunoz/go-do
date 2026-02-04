package main

import "github.com/charmbracelet/lipgloss"

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#cba6f7")).
			MarginBottom(1)

	cursorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#94e2d5"))

	checkboxPendingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#89dceb"))
	checkboxDoneStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1"))

	pendingTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#cdd6f4"))
	doneTextStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#6c7086"))

	inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fab387"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9399b2"))
	keyStyle  = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#f5c2e7"))

	doneDateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8"))

	checkmarkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1"))
	greenStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#94e2d5"))
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8"))

	containerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#45475a")).
			Padding(1)
)
