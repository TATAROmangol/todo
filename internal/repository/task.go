package repository

import (
	"context"
	"database/sql"
	"todo/internal/entities"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct{
	ctx context.Context
	db *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) *Repository{
	return &Repository{ctx, db}
}

func (r *Repository) Get() ([]entities.Task, error) {
	stmt, err := r.db.Prepare(`
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

func (r *Repository) Create(name string) (entities.Task, error) {
	stmt, err := r.db.Prepare(`
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
    row := r.db.QueryRow("SELECT LAST_INSERT_ROWID()")
    if err := row.Scan(&task.Id); err != nil {
        return entities.Task{}, err
    }

    task.Name = name
    return task, nil
}

func (r *Repository) Remove(id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM BOOKS WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
