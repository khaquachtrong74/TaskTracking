package data

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTasksToJson(tasks []TaskModel, filename string) error{
	tasksList := GetAllTasks()
	jsonData, err := json.MarshalIndent(tasksList,""," ")
	if err != nil{
		log.Fatal("Error MarshalIndent SaveTasksToJson ", err)
		return err 
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil{
		log.Fatal("Error write file SaveTasksToJson ", err)
		return err
	}
	return nil
}
