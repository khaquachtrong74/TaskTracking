package ui

import (
	"fmt"
)


func (l ListModel) View() string{

	if l.state == InputMode{
		return fmt.Sprintf(
			"Thêm công việc mới:\n\n%s\n\n(Esc để hủy, Enter để lưu)",
			l.textInput.View(),
		)
	}

	s:= "Nhiệm vụ ngày hôm nay!\n\n"
	for i, Tasks := range l.tasks{
		cursor := " "
		if l.cursor == i{
			cursor = ">"
		}
		checked := " "
		if _, ok := l.selected[i]; ok{
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s \n", cursor, checked, Tasks.Task)
	}
		s += "\n[n] Thêm mới * [q] Thoát\n"
		return s

}
