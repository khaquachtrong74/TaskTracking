package main

import (
	"GopherCron/internal/data"
	"GopherCron/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main(){
	data.InitDB()
	tasks := data.GetTasksToDay()
	initialModel := ui.InitialListModel(tasks)
	p := tea.NewProgram(
		initialModel,
	)
	if err := p.Start(); err != nil {
		panic(err)
	}
	 
}
