package models

import "github.com/edgarrps/crud-go/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	}
	return
}