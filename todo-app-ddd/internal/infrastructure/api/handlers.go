package api

import (
	"encoding/json"
	"net/http"
	"time"
	"todo-app/internal/application/commands"
	"todo-app/internal/application/queries"
)

type TodoHandler struct {
	createTodoHandler   *commands.CreateTodoHandler
	completeTodoHandler *commands.CompleteTodoHandler
	getTodosHandler     *queries.GetTodosHandler
}

func NewTodoHandler(createTodoHandler *commands.CreateTodoHandler, completeTodoHandler *commands.CompleteTodoHandler, getTodosHandler *queries.GetTodosHandler) *TodoHandler {
	return &TodoHandler{
		createTodoHandler:   createTodoHandler,
		completeTodoHandler: completeTodoHandler,
		getTodosHandler:     getTodosHandler,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var cmd commands.CreateTodoCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.createTodoHandler.Handle(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	cmd := commands.CompleteTodoCommand{ID: idStr}
	err := h.completeTodoHandler.Handle(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	query := queries.GetTodosQuery{}
	todoList, err := h.getTodosHandler.Handle(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert domain TodoList to a slice of simple structs for JSON serialization
	type simpleTodo struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
		CreatedAt string `json:"created_at"`
	}

	simpleTodos := make([]simpleTodo, 0, len(todoList.GetAllTodos()))
	for _, todo := range todoList.GetAllTodos() {
		simpleTodos = append(simpleTodos, simpleTodo{
			ID:        todo.ID.String(),
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt.Format(time.RFC3339),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(simpleTodos)
}
