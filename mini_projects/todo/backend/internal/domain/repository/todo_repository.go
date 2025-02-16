package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"to-do/internal/domain/models"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *models.Todo) error
	GetByID(ctx context.Context, id int64) (*models.Todo, error)
	GetAll(ctx context.Context) ([]*models.Todo, error)
	Update(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id int64) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, todo *models.Todo) error {
	query := `INSERT INTO todo (title)
			values ($1) RETURNING id;`

	return r.db.QueryRowContext(ctx, query, todo.Title).Scan(&todo.ID)
}

func (r *todoRepository) GetByID(ctx context.Context, id int64) (*models.Todo, error) {
	var todo models.Todo

	query := `SELECT id, title FROM todo WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&todo.ID, &todo.Title)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) GetAll(ctx context.Context) ([]*models.Todo, error) {
	query := `SELECT id,title FROM todo`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var todos []*models.Todo

	for rows.Next() {
		var todo = new(models.Todo)

		if err = rows.Scan(&todo.ID, &todo.Title); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func (r *todoRepository) Update(ctx context.Context, todo *models.Todo) error {
	query := `UPDATE todo SET title  = $1, updated_at = now() WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, todo.Title, todo.ID)

	return err
}

func (r *todoRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM todo WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
