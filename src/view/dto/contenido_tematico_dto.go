package dto

type ConteniodoTematicoDTO struct {
	ID           int64               `json:"id,omitempty"`
	Orden        int                 `json:"orden,omitempty"`
	Descripcion  string              `json:"descripcion,omitempty"`
	CursoPeriodo ItemCursoPeriodoDTO `json:"curso_periodo,omitempty"`
}

type ItemConteniodoTematicoDTO struct {
	ID          int64  `json:"contenido_tematico_id,omitempty"`
	Orden       int    `json:"orden,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

type GuardarContenidoTematicoDTO struct {
	Descripcion    string `json:"descripcion"`
	Orden          int    `json:"orden"`
	CursoPeriodoID int64  `json:"curso_periodo_id"`
}
