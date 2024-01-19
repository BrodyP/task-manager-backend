package usecase

import (
	"encoding/json"
)

type TaskData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}

func (u *usecaseImpl) getAllTask() ([]TaskData, error) {

	var result []TaskData

	jsonData, err := u.repository.GetAllTask()
	if err != nil {
		return []TaskData{}, err
	}

	for key, value := range jsonData {
		v := map[string]interface{}{}
		bytes, _ := json.Marshal(value)
		json.Unmarshal(bytes, &v)

		result = append(result, TaskData{
			Id:   key,
			Name: v["name"].(string),
			Time: v["time"].(string),
		})

	}

	return result, nil
}
