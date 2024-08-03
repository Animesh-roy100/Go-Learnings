package main

import (
	"log"
	"net/http"
	"os"
	"todo-app/internal/application/commands"
	"todo-app/internal/application/queries"
	"todo-app/internal/infrastructure/api"
	"todo-app/internal/infrastructure/persistence"
	todoHttp "todo-app/internal/interfaces/http"
)

func main() {
	// MongoDB connection string
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:password@localhost:27017"
	}

	// Initialize repository
	repo, err := persistence.NewMongoRepository(mongoURI, "todo_db", "todos")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize handlers
	createTodoHandler := commands.NewCreateTodoHandler(repo)
	completeTodoHandler := commands.NewCompleteTodoHandler(repo)
	getTodosHandler := queries.NewGetTodosHandler(repo)

	// Initialize API handler
	todoHandler := api.NewTodoHandler(createTodoHandler, completeTodoHandler, getTodosHandler)

	// Initialize router
	router := todoHttp.NewRouter(todoHandler)

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
