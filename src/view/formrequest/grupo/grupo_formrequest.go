package formrequest

type GrupoFormRequest struct {
	CelebracionID  int64   `json:"celebracion_id" binding:"required"`
	CursoPeriodoID int64   `json:"curso_periodo_id" binding:"required"`
	Maestros       []int64 `json:"maestros" binding:"required"`
}
