package ui

import (
	"GopherCron/internal/data"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)



func (l *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){

	switch l.state{
		case InputMode:
			return l.handleCaseInputMode(msg)
					
		case ViewMode:
			return l.handleCaseViewMode(msg)
				
	}
	return l, nil
}


func (l *ListModel) handleEnterInput() (tea.Model, tea.Cmd){
	if l.textInput.Value() == ""{
		return l, nil		
	}else{newTask := data.NewTask(l.textInput.Value()) 
		l.tasks = append(l.tasks, newTask)
		l.textInput.SetValue("")
		l.state = ViewMode
		data.InsertTask(newTask)
		return l, nil
	}
}
func (l *ListModel) handleEscInput()(tea.Model, tea.Cmd){
	l.textInput.SetValue("")
	l.state = ViewMode 
	return l, nil
}

func (l *ListModel) handleCaseInputMode(msg tea.Msg)(tea.Model, tea.Cmd){
	var cmd tea.Cmd
	switch m := msg.(type){
		case tea.KeyMsg:
			switch m.String(){
			case "enter":
				return l.handleEnterInput();
			case "esc":
				return l.handleEscInput();
			}
	}
	l.textInput, cmd = l.textInput.Update(msg)
	return l, cmd
}



func (l *ListModel) handleQuit()(tea.Model, tea.Cmd){
		err := data.SaveTasksToJson("tasks.json")
		if err != nil {
			fmt.Println("Lỗi lưu file: ", err)
		}
		return l, tea.Quit
}

func (l *ListModel) handleEnterTaskView()(tea.Model, tea.Cmd){
		_, ok := l.selected[l.cursor]
		if ok{
			delete(l.selected, l.cursor)
			l.tasks[l.cursor].Status = false;
		}else{
			l.selected[l.cursor] = struct{}{}
			l.tasks[l.cursor].Status = true;
		}
		data.UpdateTask(l.tasks[l.cursor]);
		return l, nil
}

func (l *ListModel) handleCaseViewMode(msg tea.Msg)(tea.Model, tea.Cmd){

	switch m := msg.(type){
	case tea.KeyMsg:
	switch m.String(){
		case "q":
			return l.handleQuit()
		case "n":
			l.state = InputMode
			l.textInput.Focus()
			return l, textinput.Blink
		case "up", "k":
			if l.cursor > 0{
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.tasks) - 1{
				l.cursor++
			}
		case "enter":
			return l.handleEnterTaskView();
		}
	}
	return l, nil  
}
