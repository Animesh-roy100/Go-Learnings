package queries

import "todo-app/internal/domain"

type GetTodosQuery struct{}

type GetTodosHandler struct {
	repo domain.TodoRepository
}

func NewGetTodosHandler(repo domain.TodoRepository) *GetTodosHandler {
	return &GetTodosHandler{repo: repo}
}

func (h *GetTodosHandler) Handle(query GetTodosQuery) (*domain.TodoList, error) {
	return h.repo.FindAll()
}
