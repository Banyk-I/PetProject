package response

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorMessage struct {
	ErrorMessage string `json:"error"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	logrus.Error(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	responseJson := ErrorMessage{
		ErrorMessage: fmt.Sprintf("%v", message),
	}

	if err := json.NewEncoder(w).Encode(responseJson); err != nil {
		http.Error(w, `{"error": "Error encoding JSON"}`, http.StatusInternalServerError)
		return
	}
}
