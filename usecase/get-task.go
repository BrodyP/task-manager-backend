package usecase

import (
	"encoding/json"
	"task-manager/model"
)

func (u *usecaseImpl) getAllTask() ([]model.TaskData, error) {
	jsonData, err := u.repository.GetAllTask()
	if err != nil {
		return nil, err
	}

	result := make([]model.TaskData, 0, len(jsonData))

	for key, value := range jsonData {
		v := make(map[string]interface{})
		bytes, _ := json.Marshal(value)
		json.Unmarshal(bytes, &v)
		result = append(result, model.TaskData{
			Id:       key,
			Name:     v["name"].(string),
			Time:     v["time"].(string),
			Complete: v["complete"].(bool),
		})
	}

	return result, nil
}
