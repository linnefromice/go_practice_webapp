package models

import "testing"

func TestTask(t *testing.T) {
	title := "title"
	description := "description"
	task := NewTask(title, description)

	t.Run(".NewTask", func(t *testing.T) {
		if task.id != 999 {
			t.Error("id is wrong")
		}
		if task.title != title {
			t.Error("title is wrong")
		}
		if task.description != description {
			t.Error("title is wrong")
		}
		if task.isFinished != false {
			t.Error("isFinished is wrong")
		}
		if task.version != 1 {
			t.Error("version is wrong")
		}
	})
	t.Run(".finished", func(t *testing.T) {
		title := "title"
		content := "content"
		task := NewTask(title, content)
		task.finished()
		if task.isFinished != true {
			t.Error("isFinished is wrong")
		}
	})
}
