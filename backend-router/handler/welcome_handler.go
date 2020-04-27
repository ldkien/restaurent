package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restaurant/backend-base/logger"
	"restaurant/backend-entity/entities"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	var baseRequest entities.BaseRequest
	err := json.NewDecoder(r.Body).Decode(&baseRequest)
	if err != nil {
		logger.Logger.Error(err)
	}
	logger.Logger.Info(baseRequest)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(fmt.Sprintf("Welcome with method %s","GET")))
	case http.MethodPost:
		w.Write([]byte(fmt.Sprintf("Welcome with method %s", "POST")))
	default:
		w.Write([]byte(fmt.Sprintf("Welcome with %s method", "other")))
	}
}
