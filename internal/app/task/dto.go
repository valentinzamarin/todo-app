package task

type TaskDTO struct {
	ID    int
	Title string
}

type GetTasksResponse struct {
	Tasks []TaskDTO
}

type CreateTaskResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
