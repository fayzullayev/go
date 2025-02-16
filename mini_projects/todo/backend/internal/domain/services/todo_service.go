package services

import (
	"context"
	"errors"
	"to-do/internal/domain/models"
	"to-do/internal/domain/repository"
)

type TodoService interface {
	CreateTodo(ctx context.Context, title string) (*models.Todo, error)
	GetTodo(ctx context.Context, id int64) (*models.Todo, error)
	ListTodos(ctx context.Context) ([]*models.Todo, error)
	UpdateTodo(ctx context.Context, todo *models.Todo) error
	DeleteTodo(ctx context.Context, id int64) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(ctx context.Context, title string) (*models.Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	todo := &models.Todo{
		Title: title,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) GetTodo(ctx context.Context, id int64) (*models.Todo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *todoService) ListTodos(ctx context.Context) ([]*models.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *todoService) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	if todo == nil {
		return errors.New("todo cannot be nil")
	}
	return s.repo.Update(ctx, todo)
}

func (s *todoService) DeleteTodo(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
