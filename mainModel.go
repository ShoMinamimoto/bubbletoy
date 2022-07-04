package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list.Model
}

func (m Model) Init() tea.Cmd {
	//TODO implement me
	panic("implement me")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}

func (m Model) View() string {
	//TODO implement me
	panic("implement me")
}

func NewModel() Model {
	return Model{}
}
