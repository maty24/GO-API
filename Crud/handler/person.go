package handler

import (
	"apis/Crud/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

// Crear una persona
func (p *person) create(w http.ResponseWriter, r *http.Request) {

	//validar que el metodo sea POST
	if r.Method != http.MethodPost { // si el metodo no es POST entonces retorna un error

		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte(`{"error": "method not allowed, basura culia"}`))

		response := newResponse(Error, "method not allowed, basura culia", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	// decodificar el body de la peticion
	data := model.Person{}                       // creo una variable data de tipo Person
	err := json.NewDecoder(r.Body).Decode(&data) // decodifico el body de la peticion y lo guardo en data, debe ser un puntero para que se guarde en la variable
	if err != nil {
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte(`{"error": "bad request, basura culia ta malo el body"}`))
		//return

		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Create(&data) // le paso la data a la funcion Create de storage
	if err != nil {
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte(`{"error": "internal server error, basura culia"}`))
		//return

		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//w.Write([]byte(`{"message": "created, alguna wea bien que hiciste"}`))

	response := newResponse(Message, "Persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, response)

}

// Obtener todas las personas
func (p *person) getAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

// Actualizar una persona
func (p *person) update(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	//el strconv.Atoi convierte un string a un entero
	ID, err := strconv.Atoi(r.URL.Query().Get("id")) //obtengo el id de la url que vienen despues del ? y lo convierto a un entero
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	responseJSON(w, http.StatusOK, response)

}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El id debe ser un número entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}
