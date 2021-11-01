package models

import (
	"errors"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsFinished  bool   `json:"isFinished"`
	IsDeleted   bool   `json:"isDeleted"`
	Version     int    `json:"version"`
}

func NewTask(id int, title, description string) Task {
	return Task{
		Id:          id,
		Title:       title,
		Description: description,
		IsFinished:  false,
		IsDeleted:   false,
		Version:     1,
	}
}

func (t *Task) Finished() {
	t.IsFinished = true
	t.Version += 1
}

func (t *Task) Deleted() {
	t.IsDeleted = true
	t.Version += 1
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
		if id == tl.tasks[i].Id {
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
	t.Title = title
	t.Description = description
	t.Version += 1
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
		if max < s.Id {
			max = s.Id
		}
	}
	return max + 1
}
