package response

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// сереалізація data in JSON і відправлення відразу через `w` і StatusOK
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error": "Error encoding JSON"}`, http.StatusInternalServerError)
		return
	}
}
