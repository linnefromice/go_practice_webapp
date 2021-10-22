package models

type Task struct {
	id          int
	title       string
	description string
	isFinished  bool
	version     int
}

func NewTask(title string, description string) Task {
	return Task{
		id:          999,
		title:       title,
		description: description,
		isFinished:  false,
		version:     1,
	}
}

func (t *Task) finished() {
	t.isFinished = true
	t.version += 1
}
