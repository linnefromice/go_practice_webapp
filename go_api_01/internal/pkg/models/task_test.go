package models

import "testing"

func TestNewTask(t *testing.T) {
	t.Run(".NewTask", func(t *testing.T) {
		title := "title"
		content := "content"
		task := NewTask(title, content)
		if task.id != 999 {
			t.Error("id is wrong")
		}
		if task.title != title {
			t.Error("title is wrong")
		}
		if task.content != content {
			t.Error("title is wrong")
		}
		if task.isFinished != false {
			t.Error("isFinished is wrong")
		}
		if task.version != 1 {
			t.Error("version is wrong")
		}
	})
}
