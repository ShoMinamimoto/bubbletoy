package main

import (
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	label       string
	description string
	timer       stopwatch.Model
}

func (t Task) FilterValue() string {
	return t.label
}

func (t Task) Init() tea.Cmd {
	return nil
}

func (t Task) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return t, tea.Batch(cmds...)
}

func (t Task) View() string {
	return t.label
}
