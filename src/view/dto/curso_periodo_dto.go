package dto

type CursoPeriodoDTO struct {
	ID         int64  `json:"curso_periodo_id"`
	PeriodoID  int64  `json:"periodo_id"`
	CursoID    int64  `json:"curso_id"`
	Nombre     string `json:"nombre"`
	EdadMinima int    `json:"edad_minima"`
	EdadMaxima int    `json:"edad_maxima"`
	Estado     string `json:"estado"`
}

type ItemCursoPeriodoDTO struct {
	ID      int64      `json:"curso_periodo_id"`
	Periodo PeriodoDTO `json:"periodo"`
	Curso   CursoDTO   `json:"curso"`
}
