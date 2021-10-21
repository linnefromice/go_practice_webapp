package models

import "testing"

func TestNewTodo(t *testing.T) {
	t.Run(".NewTodo", func(t *testing.T) {
		title := "title"
		content := "content"
		todo := NewTodo(title, content)
		if todo.id != 999 {
			t.Error("id is wrong")
		}
		if todo.title != title {
			t.Error("title is wrong")
		}
		if todo.content != content {
			t.Error("title is wrong")
		}
		if todo.isFinished != false {
			t.Error("isFinished is wrong")
		}
		if todo.version != 1 {
			t.Error("version is wrong")
		}
	})
}
