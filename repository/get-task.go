package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func (r *repositoryImpl) getAllTask() (map[string]interface{}, error) {

	var result map[string]interface{}

	resp, err := http.Get(os.Getenv("FIREBASE_URL"))

	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var jsonResp map[string]interface{}
	json.Unmarshal(b, &jsonResp)

	return jsonResp, nil

}
