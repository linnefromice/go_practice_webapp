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

func (t *TaskList) getAllTasks() []Task {
	return t.tasks
}

func (t *TaskList) getTask(index int) Task {
	return t.tasks[index]
}

func (t *TaskList) addTask(title string, description string) Task {
	task := NewTask(title, description)
	t.tasks = append(t.tasks, task)
	return task
}

func (t *TaskList) updateTask(index int, title, description string) Task {
	task := &t.tasks[index]
	task.title = title
	task.description = description
	return *task
}
