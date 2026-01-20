package data

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite", "tasks.db")

	if err != nil{
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil{
		log.Fatal(err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		status BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = DB.Exec(query)
	if err != nil{
		log.Fatal("Lỗi tạo bảng: ", err)
	}
}
