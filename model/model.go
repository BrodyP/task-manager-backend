package model

type TaskData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Time     string `json:"time"`
	Complete bool   `json:"complete"`
}

type CreateTaskData struct {
	Name     string `json:"name"`
	Time     string `json:"time"`
	Complete bool   `json:"complete"`
}

type CreateTaskResponse struct {
	IsSuccess bool      `json:"isSuccess"`
	TaskData  *TaskData `json:"taskData"`
	Error     *string   `json:"error"`
}
