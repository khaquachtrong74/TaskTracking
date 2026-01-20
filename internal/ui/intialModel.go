package ui

import (
	"GopherCron/internal/data"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (l ListModel) Init() tea.Cmd{
	return nil
}

func InitialListModel(tasks []data.TaskModel) *ListModel {
	ti := textinput.New()
	ti.Placeholder = "Nhập công việc hôm nay bạn nhé!"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return &ListModel{
		tasks:    tasks,
		cursor:   0,
		selected: make(map[int]struct{}),

		textInput: ti,
		state: ViewMode,
	}
}

