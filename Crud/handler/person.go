package handler

import (
	"apis/Crud/model"
	"encoding/json"
	"net/http"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {

	//validar que el metodo sea POST
	if r.Method != http.MethodPost { // si el metodo no es POST entonces retorna un error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "method not allowed, basura culia"}`))
		return
	}

	// decodificar el body de la peticion
	data := model.Person{}                       // creo una variable data de tipo Person
	err := json.NewDecoder(r.Body).Decode(&data) // decodifico el body de la peticion y lo guardo en data, debe ser un puntero para que se guarde en la variable
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request, basura culia ta malo el body"}`))
		return
	}

	err = p.storage.Create(&data) // le paso la data a la funcion Create de storage
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error, basura culia"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "created, alguna wea bien que hiciste"}`))

}
