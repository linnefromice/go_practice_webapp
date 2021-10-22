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
		task.finished()
		if task.isFinished != true {
			t.Error("isFinished is wrong")
		}
		if task.version != 2 {
			t.Error("version is wrong")
		}
	})
}

func TestTaskList(t *testing.T) {
	taskList := NewTaskList()

	t.Run(".NewTaskList", func(t *testing.T) {
		if len(taskList.tasks) != 0 {
			t.Error("method NewTaskList is wrong (tasks length is not zero)")
		}
	})
	t.Run(".addTask", func(t *testing.T) {
		title := "title"
		description := "description"

		taskList.addTask(title, description)
		if len(taskList.tasks) != 1 {
			t.Error("method NewTaskList is wrong")
		}
	})
}
