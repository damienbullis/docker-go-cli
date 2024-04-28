package utils

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func InitializeLog(verbose bool) {
	log.SetReportTimestamp(false)
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN!").
		Padding(0, 1).
		Background(lipgloss.Color("208")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO!").
		Padding(0, 1).
		Background(lipgloss.Color("72")).
		Foreground(lipgloss.Color("15")).
		Bold(true)

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 1).
		Background(lipgloss.Color("238")).
		Foreground(lipgloss.Color("0")).
		Bold(true)

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL").
		Padding(0, 1).
		Background(lipgloss.Color("196")).
		Foreground(lipgloss.Color("15")).
		Bold(true)

	log.SetStyles(styles)

	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Verbose output enabled")
	}

}
