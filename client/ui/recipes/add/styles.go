package add

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Base             lipgloss.Style
	Header           lipgloss.Style
	GroupHeader      lipgloss.Style
	GroupHeaderFocus lipgloss.Style
	Group            lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}

	s.Base = lg.NewStyle().
		Padding(2, 2, 0, 2)
	s.Header = lg.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Transform(strings.ToUpper).
		Align(lipgloss.Left).
		Bold(true).
		Padding(0, 2, 1, 2)
	s.GroupHeader = lg.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7861DA")).
		Align(lipgloss.Center).
		Padding(1, 2, 1, 2)
	s.GroupHeaderFocus = lg.NewStyle().
		Foreground(lipgloss.Color("#FB73C9")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7861DA")).
		Align(lipgloss.Center).
		Bold(true).
		Padding(1, 2, 1, 2)
	s.Group = lg.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Align(lipgloss.Left)

	return &s
}
