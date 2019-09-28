package errors

import (
	"2019_2_Shtoby_shto/src/utils"
	"encoding/json"
	"log"
	"net/http"
)

// Описание структуры ответа при ошибке
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ErrorHandler(w http.ResponseWriter, message string, status int, err error) {
	errorMessage := "Error!"
	if err != nil {
		errorMessage = err.Error()
	}
	b, _ := json.Marshal(&ErrorResponse{
		Status:  status,
		Message: message,
		Error:   errorMessage,
	})
	log.Println(message)
	utils.SetHeaders(&w)
	w.WriteHeader(status)
	w.Write([]byte(b))
}
