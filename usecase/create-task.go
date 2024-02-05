package usecase

import "task-manager/model"

func (u *usecaseImpl) createTask(task model.CreateTaskData) (*model.TaskData, error) {
	id, err := u.repository.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return &model.TaskData{
		Id:       id,
		Name:     task.Name,
		Time:     task.Time,
		Complete: task.Complete,
	}, nil
}
