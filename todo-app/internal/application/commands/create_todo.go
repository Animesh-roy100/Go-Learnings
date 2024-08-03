package commands

import "todo-app/internal/domain"

type CreateTodoCommand struct {
	Title string
}

type CreateTodoHandler struct {
	repo domain.TodoRepository
}

func NewCreateTodoHandler(repo domain.TodoRepository) *CreateTodoHandler {
	return &CreateTodoHandler{repo: repo}
}

func (h *CreateTodoHandler) Handle(cmd CreateTodoCommand) error {
	todo := domain.NewTodo(cmd.Title)
	return h.repo.Save(todo)
}
