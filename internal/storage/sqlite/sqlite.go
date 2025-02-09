package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"todo/internal/models"
)

type SQLite struct {
	db *sql.DB
}

func New(path string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY,
			name TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	if _, err = stmt.Exec(); err != nil {
		return nil, err
	}

	return &SQLite{db}, nil
}

func (s *SQLite) GetTasks() ([]models.Task, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, name 
		  FROM tasks
	`)
	if err != nil {
		return nil, err
	}

	var res []models.Task
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.Name)
		res = append(res, task)
	}

	return res, nil
}

func (s *SQLite) AddTask(task models.Task) error {
	stmt, err := s.db.Prepare(`INSERT INTO TASKS(id, name)
	VALUES (?,?)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(task.Id, task.Name); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) RemoveTask(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM BOOKS WHERE id = ?`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
