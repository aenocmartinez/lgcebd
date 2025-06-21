package dto

type AlumnoDTO struct {
	ID                int64  `json:"id"`
	Nombre            string `json:"nombre"`
	FechaNacimiento   string `json:"fecha_nacimiento"`
	Telefono          string `json:"telefono"`
	Acudiente         string `json:"acudiente"`
	AcudienteTelefono string `json:"acudiente_telefono"`
	Direccion         string `json:"direccion"`
	Edad              int    `json:"edad,omitempty"`
	Activo            bool   `json:"activo"`
}
