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

func (t *Task) finished() {
	t.isFinished = true
	t.version += 1
}

func (t *Task) deleted() {
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

func (tl *TaskList) getAllTasks() []Task {
	return tl.tasks
}

func (tl *TaskList) getTask(id int) (*Task, error) {
	for i := 0; i < len(tl.tasks); i++ {
		if id == tl.tasks[i].id {
			return &tl.tasks[i], nil
		}
	}

	return &Task{}, errors.New("task is not found")
}

func (tl *TaskList) addTask(title string, description string) Task {
	nextId := tl.nextId()
	t := NewTask(nextId, title, description)
	tl.tasks = append(tl.tasks, t)
	return t
}

func (tl *TaskList) updateTask(id int, title, description string) Task {
	t, _ := tl.getTask(id)
	t.title = title
	t.description = description
	t.version += 1
	return *t
}

func (tl *TaskList) finishTask(id int) Task {
	t, _ := tl.getTask(id)
	t.finished()
	return *t
}

func (tl *TaskList) deleteTask(id int) Task {
	t, _ := tl.getTask(id)
	t.deleted()
	return *t
}

func (tl *TaskList) nextId() int {
	max := 0
	for _, s := range tl.tasks {
		if max < s.id {
			max = s.id
		}
	}
	return max + 1
}
