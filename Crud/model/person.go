package model

// Community estructura de una comunidad
type Community struct {
	// Name nombre de una comunidad.
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []Community

// Person estructura de una persona
type Person struct {
	// Name nombre de la persona Ej: Matias
	Name string `json:"name"`
	// Age edad de la persona Ej: 26
	Age uint8 `json:"age"`
	// Communities comunidades a las que pertenece una persona
	Communities Communities `json:"communities"`
}

// Persons slice de personas
type Persons []Person
