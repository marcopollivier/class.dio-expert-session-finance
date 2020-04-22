package actuator

import (
	"encoding/json"
	"net/http"
)

func Health(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	profile := HealthBody{"alive"}

	returnBody, err := json.Marshal(profile)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = responseWriter.Write(returnBody)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

type HealthBody struct {
	Status string `json:"status"`
}
