package errorGen

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/models/Error"
	"log"
	"net/http"
)

func ErrorGen(w http.ResponseWriter, message string, code int) {
	var err Error.Error
	err.Message = message
	log.Println("ИДИСЮДА:", err)
	w.WriteHeader(code)

	// Затем кодируем и отправляем JSON с сообщением об ошибке
	if encodeErr := json.NewEncoder(w).Encode(err); encodeErr != nil {
		log.Println("Error encoding JSON:", encodeErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
