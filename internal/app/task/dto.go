package task

type TaskDTO struct {
	ID    int
	Title string
}

type GetTasksResponse struct {
	Tasks []TaskDTO
}
