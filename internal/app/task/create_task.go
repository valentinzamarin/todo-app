package task

import (
	"errors"
	"strings"
	"todo-app/internal/domain/task"
)

type CreateTask struct {
	taskRepo task.TaskRepo
}

func NewCreateTask(taskRepo task.TaskRepo) *CreateTask {
	return &CreateTask{
		taskRepo: taskRepo,
	}
}

func (ct *CreateTask) Execute(title string) (*CreateTaskResponse, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title cannot be empty")
	}

	createdTask, err := ct.taskRepo.CreateTaskCallback(title)
	if err != nil {
		return nil, err
	}

	return &CreateTaskResponse{
		ID:    createdTask.ID,
		Title: createdTask.Title,
	}, nil
}
