package models

type Task struct {
	id          int
	title       string
	description string
	isFinished  bool
	isDeleted   bool
	version     int
}

func NewTask(title string, description string) Task {
	return Task{
		id:          999,
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

func (tl *TaskList) getTask(index int) Task {
	return tl.tasks[index]
}

func (tl *TaskList) addTask(title string, description string) Task {
	t := NewTask(title, description)
	tl.tasks = append(tl.tasks, t)
	return t
}

func (tl *TaskList) updateTask(index int, title, description string) Task {
	t := &tl.tasks[index]
	t.title = title
	t.description = description
	return *t
}
