package models

import (
	"errors"
)

type Task struct {
	id          int
	title       string
	description string
	isFinished  bool
	isDeleted   bool
	version     int
}

func NewTask(id int, title, description string) Task {
	return Task{
		id:          id,
		title:       title,
		description: description,
		isFinished:  false,
		isDeleted:   false,
		version:     1,
	}
}

func (t *Task) Finished() {
	t.isFinished = true
	t.version += 1
}

func (t *Task) Deleted() {
	t.isDeleted = true
	t.version += 1
}

type TaskList struct {
	tasks []Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		[]Task{},
	}
}

func (tl *TaskList) GetAllTasks() []Task {
	return tl.tasks
}

func (tl *TaskList) GetTask(id int) (*Task, error) {
	for i := 0; i < len(tl.tasks); i++ {
		if id == tl.tasks[i].id {
			return &tl.tasks[i], nil
		}
	}

	return &Task{}, errors.New("task is not found")
}

func (tl *TaskList) AddTask(title string, description string) Task {
	nextId := tl.NextId()
	t := NewTask(nextId, title, description)
	tl.tasks = append(tl.tasks, t)
	return t
}

func (tl *TaskList) UpdateTask(id int, title, description string) Task {
	t, _ := tl.GetTask(id)
	t.title = title
	t.description = description
	t.version += 1
	return *t
}

func (tl *TaskList) FinishTask(id int) Task {
	t, _ := tl.GetTask(id)
	t.Finished()
	return *t
}

func (tl *TaskList) DeleteTask(id int) Task {
	t, _ := tl.GetTask(id)
	t.Deleted()
	return *t
}

func (tl *TaskList) NextId() int {
	max := 0
	for _, s := range tl.tasks {
		if max < s.id {
			max = s.id
		}
	}
	return max + 1
}
