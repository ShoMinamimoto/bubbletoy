package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	program := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if err := program.Start(); err != nil {
		fmt.Printf("Something went wrong: %v", err)
		log.Fatal(err)
	}
}
