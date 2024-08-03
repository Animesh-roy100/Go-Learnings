package domain

import "errors"

type TodoList struct {
	todos []*Todo
}

func NewTodoList() *TodoList {
	return &TodoList{todos: []*Todo{}}
}

func (tl *TodoList) AddTodo(todo *Todo) {
	tl.todos = append(tl.todos, todo)
}

func (tl *TodoList) CompleteTodo(id TodoID) error {
	for _, todo := range tl.todos {
		if todo.ID == id {
			todo.Complete()
			return nil
		}
	}
	return errors.New("todo not found")
}

func (tl *TodoList) GetAllTodos() []*Todo {
	return tl.todos
}

type TodoRepository interface {
	Save(todo *Todo) error
	FindAll() (*TodoList, error)
	Update(todo *Todo) error
}
