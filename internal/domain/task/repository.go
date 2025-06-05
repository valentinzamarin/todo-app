package task

type TaskRepo interface {
	GetAll() ([]Task, error)
	CreateTaskCallback(title string) (*Task, error)
}
