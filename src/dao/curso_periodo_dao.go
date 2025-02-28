package dao

import (
	"ebd/src/domain"
	"ebd/src/view/dto"

	"gorm.io/gorm"
)

type CursoPeriodoDao struct {
	db *gorm.DB
}

func NewCursoPeriodoDao(db *gorm.DB) *CursoPeriodoDao {
	return &CursoPeriodoDao{db: db}
}

type cursoPeriodoData struct {
	ID        int64  `gorm:"column:id"`
	PeriodoID int64  `gorm:"column:periodo_id"`
	CursoID   int64  `gorm:"column:curso_id"`
	Nombre    string `gorm:"column:nombre"`
	EdadMin   int    `gorm:"column:edad_minima"`
	EdadMax   int    `gorm:"column:edad_maxima"`
}

func (r *CursoPeriodoDao) FindByPeriodoYCurso(periodoID, cursoID int64) (*domain.CursoPeriodo, error) {
	var data cursoPeriodoData
	result := r.db.Table("periodo_cursos").Where("periodo_id = ? AND curso_id = ?", periodoID, cursoID).First(&data)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewCursoPeriodoEmpty(), nil
		}
		return nil, result.Error
	}

	curso := domain.NewCurso(nil)
	curso.SetID(data.CursoID)

	periodo := domain.NewPeriodo(nil)
	periodo.SetID(data.PeriodoID)

	return domain.NewCursoPeriodo(data.ID, curso, periodo), nil
}

func (r *CursoPeriodoDao) ObtenerCursosPorPeriodo(periodoID int64) ([]dto.CursoPeriodoDTO, error) {
	var cursosData []cursoPeriodoData

	query := `
		SELECT pc.id, pc.periodo_id, pc.curso_id, c.nombre, c.edad_minima, c.edad_maxima
		FROM periodo_cursos pc
		INNER JOIN cursos c ON c.id = pc.curso_id
		WHERE pc.periodo_id = ?
	`
	result := r.db.Raw(query, periodoID).Scan(&cursosData)
	if result.Error != nil {
		return nil, result.Error
	}

	var cursosDTO []dto.CursoPeriodoDTO
	for _, c := range cursosData {
		cursosDTO = append(cursosDTO, dto.CursoPeriodoDTO{
			ID:         c.ID,
			PeriodoID:  c.PeriodoID,
			CursoID:    c.CursoID,
			Nombre:     c.Nombre,
			EdadMinima: c.EdadMin,
			EdadMaxima: c.EdadMax,
		})
	}

	return cursosDTO, nil
}

func (r *CursoPeriodoDao) FindByID(id int64) *domain.CursoPeriodo {
	var data struct {
		ID            int64  `gorm:"column:id"`
		CursoID       int64  `gorm:"column:curso_id"`
		CursoNombre   string `gorm:"column:curso_nombre"`
		EdadMinima    int    `gorm:"column:edad_minima"`
		EdadMaxima    int    `gorm:"column:edad_maxima"`
		PeriodoID     int64  `gorm:"column:periodo_id"`
		PeriodoNombre string `gorm:"column:periodo_nombre"`
		FechaInicio   string `gorm:"column:fecha_inicio"`
		FechaFin      string `gorm:"column:fecha_fin"`
	}

	query := `
		SELECT 
			pc.id, pc.periodo_id, pc.curso_id,
			c.nombre AS curso_nombre, c.edad_minima, c.edad_maxima,
			p.nombre AS periodo_nombre, p.fecha_inicio, p.fecha_fin
		FROM periodo_cursos pc
		INNER JOIN cursos c ON c.id = pc.curso_id
		INNER JOIN periodos p ON p.id = pc.periodo_id
		WHERE pc.id = ?
	`
	result := r.db.Raw(query, id).Scan(&data)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewCursoPeriodoEmpty()
		}
		return domain.NewCursoPeriodoEmpty()
	}

	curso := domain.NewCurso(nil)
	curso.SetID(data.CursoID)
	curso.SetNombre(data.CursoNombre)
	curso.SetEdadMinima(data.EdadMinima)
	curso.SetEdadMaxima(data.EdadMaxima)

	periodo := domain.NewPeriodo(nil)
	periodo.SetID(data.PeriodoID)
	periodo.SetNombre(data.PeriodoNombre)
	periodo.SetFechaInicio(data.FechaInicio)
	periodo.SetFechaFin(data.FechaFin)

	return domain.NewCursoPeriodo(data.ID, curso, periodo)
}
