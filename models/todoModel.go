package models

import (
	"time"

	"x.com/todo/database"
)

type Todo struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      int64     `json:"user_id"`
}

func (todo *Todo) Delete() error {
	query := `DELETE FROM todos WHERE id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.Id)
	return err
}

func (todo *Todo) Update() error {
	query := `
	UPDATE todos 
	SET title = ?, description = ?, updated_at = ?
	WHERE id = ?
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, time.Now(), todo.Id)
	return err
}

func GetTodo(id int64) (*Todo, error) {
	query := `SELECT * FROM todos WHERE id = ?`

	var todo Todo

	row := database.DB.QueryRow(query, id)
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (todo *Todo) Save() error {
	query := `INSERT INTO 
	todos(id, title, description, user_id)
	VALUES(?, ?, ?, ?)
	`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(todo.Id, todo.Title, todo.Description, todo.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.Id = id
	return nil
}

func GetUsersTodos(userId int64) ([]Todo, error) {
	query := `SELECT * FROM todos WHERE user_id = ?`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if len(todos) == 0 {
		todos = []Todo{}
	}

	return todos, nil
}
