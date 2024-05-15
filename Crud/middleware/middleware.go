package middleware

import (
	"log"
	"net/http"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("peticion %q, método: %q", r.URL.Path, r.Method) //le enviamos la url y el metodo
		f(w, r)                                                     //ejecutamos la función que recibe los parametros w y r
	}
}

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization") // leemos el token del header
		//validamos el token
		if token != "token-seguroxd" {
			forbidden(w, r)
			return
		}

		f(w, r) //ejecutamos la función que recibe los parametros w y r
	}

}

// esta función se encarga de enviar un mensaje de error cuando no se tiene autorización
func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Quien chota sos no tienes autorización maquinola"))
}
