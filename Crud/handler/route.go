package handler

import "net/http"

func RoutePerson(mux *http.ServeMux, storage Storage) {
	//el http.ServeMux es un enrutador HTTP que se basa en la URL de la petición para redirigir la petición al controlador correspondiente.

	p := newPerson(storage)

	mux.HandleFunc("/v1/person/create", p.create)
	mux.HandleFunc("/v1/person/getAll", p.getAll)
	mux.HandleFunc("/v1/person/update", p.update)

}
