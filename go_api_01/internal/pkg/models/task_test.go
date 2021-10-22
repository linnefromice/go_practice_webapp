package models

import "testing"

func TestTask(t *testing.T) {
	title := "title"
	description := "description"
	var task Task

	t.Run(".NewTask", func(t *testing.T) {
		task = NewTask(title, description)
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
		if task.isFinished != false {
			t.Error("isFinished is wrong")
		}
		if task.version != 1 {
			t.Error("version is wrong")
		}
	})
	t.Run(".finished", func(t *testing.T) {
		task = NewTask(title, description)
		task.finished()
		if task.isFinished != true {
			t.Error("isFinished is wrong")
		}
		if task.version != 2 {
			t.Error("version is wrong")
		}
	})
	t.Run(".deleted", func(t *testing.T) {
		task = NewTask(title, description)
		task.deleted()
		if task.isDeleted != true {
			t.Error("isDeleted is wrong")
		}
		if task.version != 2 {
			t.Error("version is wrong")
		}
	})
}

func TestTaskList(t *testing.T) {
	var taskList *TaskList
	t.Run(".NewTaskList", func(t *testing.T) {
		taskList = NewTaskList()
		if len(taskList.tasks) != 0 {
			t.Error("method NewTaskList is wrong (tasks length is not zero)")
		}
	})
	t.Run(".getAllTasks", func(t *testing.T) {
		taskList = NewTaskList()
		if len(taskList.getAllTasks()) != 0 {
			t.Error("method NewTaskList is wrong (tasks length is not zero)")
		}
	})
	t.Run(".getTask / .addTask", func(t *testing.T) {
		taskList = NewTaskList()
		title := "title"
		description := "description"

		taskList.addTask(title, description)
		if len(taskList.tasks) != 1 {
			t.Error("method NewTaskList is wrong")
		}

		task := taskList.getTask(len(taskList.getAllTasks()) - 1)
		if task.title != title {
			t.Error("title is wrong")
		}
		if task.description != description {
			t.Error("title is wrong")
		}
	})
	t.Run(".updateTask", func(t *testing.T) {
		taskList = NewTaskList()
		taskList.addTask("initial title", "initial description")
		updatedTitle := "updated title"
		updatedDescription := "updated description"
		index := len(taskList.getAllTasks()) - 1
		taskList.updateTask(index, updatedTitle, updatedDescription)
		task := taskList.getTask(index)
		if task.title != updatedTitle {
			t.Error("title is wrong")
		}
		if task.description != updatedDescription {
			t.Error("title is wrong")
		}
	})
}
