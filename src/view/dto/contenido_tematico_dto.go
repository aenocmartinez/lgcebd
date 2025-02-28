package dto

type ConteniodoTematicoDTO struct {
	ID           int64               `json:"id"`
	Descripcion  string              `json:"descripcion"`
	CursoPeriodo ItemCursoPeriodoDTO `json:"curso_periodo"`
}

type GuardarContenidoTematicoDTO struct {
	Descripcion    string `json:"descripcion"`
	CursoPeriodoID int64  `json:"curso_periodo_id"`
}
