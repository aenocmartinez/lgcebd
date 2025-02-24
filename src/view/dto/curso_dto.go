package dto

type CursoDTO struct {
	ID         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	EdadMinima int    `json:"edad_minima"`
	EdadMaxima int    `json:"edad_maxima"`
	Estado     string `json:"estado"`
}
