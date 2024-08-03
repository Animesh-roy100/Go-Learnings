package commands

import "todo-app/internal/domain"

type CompleteTodoCommand struct {
	ID string
}

type CompleteTodoHandler struct {
	repo domain.TodoRepository
}

func NewCompleteTodoHandler(repo domain.TodoRepository) *CompleteTodoHandler {
	return &CompleteTodoHandler{repo: repo}
}

func (h *CompleteTodoHandler) Handle(cmd CompleteTodoCommand) error {
	todoList, err := h.repo.FindAll()
	if err != nil {
		return err
	}

	todoID, err := domain.TodoIDFromString(cmd.ID)
	if err != nil {
		return err
	}

	if err := todoList.CompleteTodo(todoID); err != nil {
		return err
	}

	for _, todo := range todoList.GetAllTodos() {
		if todo.ID == todoID {
			return h.repo.Update(todo)
		}
	}

	return nil
}
