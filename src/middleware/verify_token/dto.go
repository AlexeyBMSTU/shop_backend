package verify_token

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Name    string `json:"name"`
}

func toDTO(w *http.ResponseWriter, message string, name string) {
	var a Response

	a.Message = message
	a.Status = http.StatusOK
	a.Name = name

	err := json.NewEncoder(*w).Encode(a)
	if err != nil {
		return
	}

	return
}
