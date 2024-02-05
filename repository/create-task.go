package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"task-manager/model"
)

const firebaseURL = "FIREBASE_URL"

func (r *repositoryImpl) createTask(m model.CreateTaskData) (string, error) {

	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv(firebaseURL), bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("unexpected status code: " + resp.Status)
	}

	var jsonResp map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&jsonResp); err != nil {
		return "", err
	}

	name, exists := jsonResp["name"].(string)
	if !exists {
		return "", errors.New("No name in the response")
	}

	return name, nil
}
