package main

import (
	"apis/Crud/handler"
	"apis/Crud/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux() //el mux es un enrutador HTTP que se basa en la URL de la petición para redirigir la petición al controlador correspondiente.

	handler.RoutePerson(mux, &store)

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error: %v", err)
	}

}
