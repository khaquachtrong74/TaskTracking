package data

import (
	"log"
	"strings"
	"time"
)
type TaskModel struct {
	ID int `json:"id"`
	Task string `json:"task"`
	Status bool		 `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTask(task string) TaskModel{
	return TaskModel{
		Task: strings.TrimSpace(task),
		Status: false,
		CreatedAt: time.Now(),
	}
}

func InsertTask(task TaskModel) (bool, error){
	query := "INSERT INTO tasks (task, status, created_at) VALUES(?, ?, ?)"
	_, err := DB.Exec(query, task.Task, task.Status, task.CreatedAt)
	if err != nil{
		log.Println("Lỗi insert task: ", err)
		return false, err 
	}
	return true, nil
}

func GetTasks(conditions string) []TaskModel{
	 query := "SELECT id, task, status, created_at FROM tasks "+conditions
	rows, err := DB.Query(query)
	if err != nil{
		log.Fatal("Lỗi Get Tasks: ", err)
	}
	defer rows.Close()
	var tasks []TaskModel
	for rows.Next(){
		var t TaskModel
		err := rows.Scan(&t.ID, &t.Task, &t.Status, &t.CreatedAt)
		if err != nil{
			log.Fatal("Lỗi scan: ", err)
			continue 
		}
		tasks = append(tasks, t)
	}
	return tasks 
}
func GetAllTasks()[]TaskModel{
	return GetTasks("")
}
func GetTasksToDay()[]TaskModel{
	today := time.Now().Format("2006-01-02")
	return GetTasks("WHERE created_at LIKE '" + today + "%' AND status=false")
}

func UpdateTask(t TaskModel){
	query := "UPDATE tasks SET status = ? WHERE id = ?"
	_, err := DB.Exec(query, t.Status, t.ID)
	if err != nil{
		log.Fatal("Lỗi update task: ", err)
	}
}
