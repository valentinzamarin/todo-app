package handlers

import (
	"encoding/json"
	"net/http"
	"todo-app/internal/app/task"
	"todo-app/internal/interfaces/dto"
)

type TaskHandler struct {
	getTasksUseCase *task.GetTasksUsecase
}

func NewTaskHandler(getTasksUsecase *task.GetTasksUsecase) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase: getTasksUsecase,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasksResponse, err := h.getTasksUseCase.Execute()
	if err != nil {
		// fmt.Println("Пусто ")
		http.Error(w, "Ошибка получения задач", http.StatusInternalServerError)
		return
	}

	response := make([]dto.TaskResponse, len(tasksResponse.Tasks))
	for i, t := range tasksResponse.Tasks {
		response[i] = dto.TaskResponse{
			ID:    t.ID,
			Title: t.Title,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
