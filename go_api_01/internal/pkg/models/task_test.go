package models

import "testing"

func TestTask(t *testing.T) {
	id := 999
	title := "title"
	description := "description"
	var task Task

	t.Run(".NewTask", func(t *testing.T) {
		task = NewTask(id, title, description)
		if task.id != id {
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
		task = NewTask(id, title, description)
		task.Finished()
		if task.isFinished != true {
			t.Error("isFinished is wrong")
		}
		if task.version != 2 {
			t.Error("version is wrong")
		}
	})
	t.Run(".deleted", func(t *testing.T) {
		task = NewTask(id, title, description)
		task.Deleted()
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
		if len(taskList.GetAllTasks()) != 0 {
			t.Error("method NewTaskList is wrong (tasks length is not zero)")
		}
	})
	t.Run(".getTask / .addTask", func(t *testing.T) {
		taskList = NewTaskList()
		title := "title"
		description := "description"

		taskList.AddTask(title, description)
		if len(taskList.tasks) != 1 {
			t.Error("method NewTaskList is wrong")
		}

		task, _ := taskList.GetTask(1)
		if task.title != title {
			t.Error("title is wrong")
		}
		if task.description != description {
			t.Error("title is wrong")
		}
	})
	t.Run(".updateTask", func(t *testing.T) {
		taskList = NewTaskList()
		taskList.AddTask("initial title", "initial description")
		updatedTitle := "updated title"
		updatedDescription := "updated description"
		index := 1
		taskList.UpdateTask(index, updatedTitle, updatedDescription)
		task, _ := taskList.GetTask(index)
		if task.title != updatedTitle {
			t.Errorf("title is wrong : %v / %v", task.title, updatedTitle)
		}
		if task.description != updatedDescription {
			t.Errorf("description is wrong : %v / %v", task.description, updatedDescription)
		}
		if task.version != 2 {
			t.Errorf("version is wrong : %v", task.version)
		}
	})
	t.Run(".finishTask", func(t *testing.T) {
		taskList = NewTaskList()
		taskList.AddTask("", "")
		index := 1

		taskList.FinishTask(index)

		task, _ := taskList.GetTask(index)
		if task.isFinished != true {
			t.Errorf("isFinished is wrong : %v", task.isFinished)
		}
		if task.version != 2 {
			t.Errorf("version is wrong : %v", task.version)
		}
	})
	t.Run(".deleteTask", func(t *testing.T) {
		taskList = NewTaskList()
		taskList.AddTask("", "")
		index := 1

		taskList.DeleteTask(index)

		task, _ := taskList.GetTask(index)
		if task.isDeleted != true {
			t.Errorf("isDeleted is wrong : %v", task.isDeleted)
		}
		if task.version != 2 {
			t.Errorf("version is wrong : %v", task.version)
		}
	})
	t.Run(".nextId (first)", func(t *testing.T) {
		taskList = NewTaskList()
		if taskList.NextId() != 1 {
			t.Error("nextId is wrong")
		}
	})
	t.Run(".nextId", func(t *testing.T) {
		taskList = NewTaskList()
		taskList.AddTask("", "")
		taskList.AddTask("", "")
		if taskList.NextId() != 3 {
			t.Error("nextId is wrong")
		}
	})
}
