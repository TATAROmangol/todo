package sqlite

import (
	"database/sql"
	"todo/internal/entities"

	_ "github.com/mattn/go-sqlite3"
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
			id INTEGER PRIMARY KEY AUTOINCREMENT,
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

func (s *SQLite) GetTasks() ([]entities.Task, error) {
	stmt, err := s.db.Prepare(`
		SELECT id, name 
		  FROM tasks
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []entities.Task
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task entities.Task
		rows.Scan(&task.Id, &task.Name)
		res = append(res, task)
	}

	return res, nil
}

func (s *SQLite) CreateTask(name string) (entities.Task, error) {
	stmt, err := s.db.Prepare(`
		INSERT INTO tasks(name)
		VALUES (?)
	`)
	if err != nil {
		return entities.Task{}, err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(name); err != nil {
        return entities.Task{}, err
    }

	var task entities.Task
    row := s.db.QueryRow("SELECT LAST_INSERT_ROWID()")
    if err := row.Scan(&task.Id); err != nil {
        return entities.Task{}, err
    }

    task.Name = name
    return task, nil
}

func (s *SQLite) RemoveTask(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM BOOKS WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
