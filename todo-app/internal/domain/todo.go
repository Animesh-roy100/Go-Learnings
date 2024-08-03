package domain

import "time"

type Todo struct {
	ID        TodoID
	Title     string
	Completed bool
	CreatedAt time.Time
}

func NewTodo(title string) *Todo {
	return &Todo{
		ID:        NewTodoID(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func (t *Todo) Complete() {
	t.Completed = true
}
