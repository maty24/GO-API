package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"` //el json es un ejemplo de como se va a ver el json que se va a retornar, se pone ese mensaje y el contenido que va a tener
	Message     string      `json:"message"`
	Data        interface{} `json:"data"` //el tipo interface{} es un tipo de dato que puede ser cualquier cosa
}

func newResponse(messageType string, message string, data interface{}) response { //esto me retorna un response
	return response{
		messageType,
		message,
		data,
	}
}

// esta funcion se encarga de responderle al cliente con un json, no retorna nada porque ya se encarga de responder
func responseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
