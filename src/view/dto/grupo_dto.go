package dto

type GrupoDto struct {
	ID           int64               `json:"id"`
	Celebracion  CelebracionDto      `json:"celebracion"`
	CursoPeriodo ItemCursoPeriodoDTO `json:"curso_periodo"`
	Maestros     []MaestroDTO        `json:"maestros"`
}

type GuardarGrupoDto struct {
	ID             int64   `json:"id"`
	CelebracionID  int64   `json:"celebracion_id"`
	CursoPeriodoID int64   `json:"curso_periodo_id"`
	Maestros       []int64 `json:"maestros"`
}
