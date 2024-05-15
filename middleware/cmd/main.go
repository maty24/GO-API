package main

import "apis/middleware/funciones"

func execute(name string, f funciones.MyFunction) {
	f(name)
}

func main() {
	name := "Comunidad mati el pulento 10"
	//los milddleeares son funciones que se agregan a una funcion principal pueder ir o no, son como los callbacks de javascript
	execute(name, funciones.MiddlewareLog(funciones.Saludar)) // MiddlewareLog recibe una funcion y retorna una funcion
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
