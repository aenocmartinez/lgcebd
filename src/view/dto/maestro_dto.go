package dto

type MaestroDTO struct {
	ID              int64  `json:"id"`
	Nombre          string `json:"nombre"`
	Telefono        string `json:"telefono"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Estado          string `json:"estado"`
}
