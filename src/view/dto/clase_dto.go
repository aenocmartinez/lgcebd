package dto

type ClaseDTO struct {
	ID                int64                 `json:"id,omitempty"`
	Fecha             string                `json:"fecha,omitempty"`
	Ofreda            float64               `json:"ofrenda,omitempty"`
	Grupo             GrupoDto              `json:"grupo,omitempty"`
	ContenidoTematico ConteniodoTematicoDTO `json:"tema,omitempty"`
}

type GuardarClaseDTO struct {
	Fecha               string  `json:"fecha"`
	Ofrenda             float64 `json:"ofrenda"`
	GrupoID             int64   `json:"grupo_id"`
	ContenidoTematicoID int64   `json:"contenido_tematico_id"`
}
