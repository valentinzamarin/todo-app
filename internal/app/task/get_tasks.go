package task

import "todo-app/internal/domain/task"

type GetTasksUsecase struct {
	taskRepo task.TaskRepo
}

func NewGetTasksUseCase(taskRepo task.TaskRepo) *GetTasksUsecase {
	return &GetTasksUsecase{
		taskRepo: taskRepo,
	}
}

func (uc *GetTasksUsecase) Execute() (*GetTasksResponse, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	dtoTasks := make([]TaskDTO, len(tasks))
	for i, t := range tasks {
		dtoTasks[i] = TaskDTO{
			ID:    t.ID,
			Title: t.Title,
		}
	}

	return &GetTasksResponse{Tasks: dtoTasks}, nil

}
