package models

type Task struct {
	id         int
	title      string
	content    string
	isFinished bool
	version    int
}

func NewTask(title string, content string) Task {
	return Task{
		id:         999,
		title:      title,
		content:    content,
		isFinished: false,
		version:    1,
	}
}

func (t *Task) finished() {
	t.isFinished = true
	t.version += 1
}
