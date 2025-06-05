package handlers

import (
	"encoding/json"
	"net/http"
	"todo-app/internal/app/task"
	"todo-app/internal/interfaces/dto"
)

type TaskHandler struct {
	getTasksUseCase   *task.GetTasksUsecase
	createTaskUseCase *task.CreateTask
}

func NewTaskHandler(getTasksUsecase *task.GetTasksUsecase, createTaskUseCase *task.CreateTask) *TaskHandler {
	return &TaskHandler{
		getTasksUseCase:   getTasksUsecase,
		createTaskUseCase: createTaskUseCase,
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

func (h *TaskHandler) CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	var req task.CreateTaskResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	createResponse, err := h.createTaskUseCase.Execute(req.Title)
	if err != nil {
		http.Error(w, "Ошибка создания задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createResponse)
}
