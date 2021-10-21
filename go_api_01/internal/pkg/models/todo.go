package models

type Todo struct {
	id         int
	title      string
	content    string
	isFinished bool
	version    int
}

func NewTodo(title string, content string) Todo {
	return Todo{
		id:         999,
		title:      title,
		content:    content,
		isFinished: false,
		version:    1,
	}
}
