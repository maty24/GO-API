package handler

import "apis/Crud/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error // me recibe un ID y una persona y me retorna un error
	Delete(ID int) error                       // me recibe un ID y me retorna un error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
