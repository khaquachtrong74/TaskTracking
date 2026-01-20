package data

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTasksToJson(filename string) error{
	tasksList := GetAllTasks()
	jsonData, err := json.MarshalIndent(tasksList,""," ")
	if err != nil{
		log.Fatal("Error MarshalIndent SaveTasksToJson ", err)
		return err 
	}
	return = os.WriteFile(filename, jsonData, 0644)
}
