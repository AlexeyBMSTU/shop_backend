package errorGen

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/models/Error"
	"log"
	"net/http"
)

func ErrorGen(w *http.ResponseWriter, message string, code int) {
	var err Error.Error
	err.Message = message
	(*w).WriteHeader(code)

	if encodeErr := json.NewEncoder(*w).Encode(err); encodeErr != nil {
		log.Println("Error encoding JSON:", encodeErr)
		http.Error(*w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
