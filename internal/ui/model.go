package ui

import (
	"GopherCron/internal/data"

	"github.com/charmbracelet/bubbles/textinput"
)

type sessionState int 
const(
	ViewMode sessionState = iota
	InputMode 
)



type ListModel struct {
	tasks [] data.TaskModel
	cursor int 
	selected map[int]struct{}

	state sessionState
	textInput textinput.Model
}
