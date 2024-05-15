package handler

import (
	"apis/Crud/middleware"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	//el http.ServeMux es un enrutador HTTP que se basa en la URL de la petición para redirigir la petición al controlador correspondiente.

	p := newPerson(storage)

	mux.HandleFunc("/v1/person/create", middleware.Log(middleware.Authentication(p.create)))
	mux.HandleFunc("/v1/person/getAll", middleware.Log(p.getAll))
	mux.HandleFunc("/v1/person/update", p.getAll)

}
