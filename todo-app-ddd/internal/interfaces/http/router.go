package http

import (
	"todo-app/internal/infrastructure/api"

	"github.com/gorilla/mux"
)

func NewRouter(todoHandler *api.TodoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	r.HandleFunc("/todos/complete", todoHandler.CompleteTodo).Methods("PUT")

	return r
}
