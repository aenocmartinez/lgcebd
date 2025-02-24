package dto

type CursoPeriodoDTO struct {
	PeriodoID  int64  `json:"periodo_id"`
	CursoID    int64  `json:"curso_id"`
	Nombre     string `json:"nombre"`
	EdadMinima int    `json:"edad_minima"`
	EdadMaxima int    `json:"edad_maxima"`
	Estado     string `json:"estado"`
}
