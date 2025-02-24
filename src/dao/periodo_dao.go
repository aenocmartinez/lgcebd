package dao

import (
	"ebd/src/domain"
	"ebd/src/view/dto"
	"log"

	"gorm.io/gorm"
)

type PeriodoDao struct {
	db *gorm.DB
}

func NewPeriodoDao(db *gorm.DB) *PeriodoDao {
	return &PeriodoDao{db: db}
}

type periodoDB struct {
	ID          int64  `gorm:"column:id"`
	Nombre      string `gorm:"column:nombre"`
	FechaInicio string `gorm:"column:fecha_inicio"`
	FechaFin    string `gorm:"column:fecha_fin"`
}

func (periodoDB) TableName() string {
	return "periodos"
}

func (r *PeriodoDao) FindByID(id int64) (*domain.Periodo, error) {
	var periodoData periodoDB
	result := r.db.Where("id = ?", id).First(&periodoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &domain.Periodo{}, nil
		}
		return nil, result.Error
	}

	periodo := domain.NewPeriodo(r)
	periodo.SetID(periodoData.ID)
	periodo.SetNombre(periodoData.Nombre)
	periodo.SetFechaInicio(periodoData.FechaInicio)
	periodo.SetFechaFin(periodoData.FechaFin)

	return periodo, nil
}

func (r *PeriodoDao) FindByNombre(nombre string) (*domain.Periodo, error) {
	var periodoData periodoDB
	result := r.db.Where("nombre = ?", nombre).First(&periodoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &domain.Periodo{}, nil
		}
		return nil, result.Error
	}

	periodo := domain.NewPeriodo(r)
	periodo.SetID(periodoData.ID)
	periodo.SetNombre(periodoData.Nombre)
	periodo.SetFechaInicio(periodoData.FechaInicio)
	periodo.SetFechaFin(periodoData.FechaFin)

	return periodo, nil
}

func (r *PeriodoDao) List() ([]dto.PeriodoDTO, error) {
	var periodosData []periodoDB
	result := r.db.Find(&periodosData)
	if result.Error != nil {
		return nil, result.Error
	}

	periodosDTO := []dto.PeriodoDTO{}
	for _, periodo := range periodosData {
		periodosDTO = append(periodosDTO, dto.PeriodoDTO{
			ID:          periodo.ID,
			Nombre:      periodo.Nombre,
			FechaInicio: periodo.FechaInicio,
			FechaFin:    periodo.FechaFin,
		})
	}

	return periodosDTO, nil
}

func (r *PeriodoDao) Save(periodo *domain.Periodo) error {
	periodoData := periodoDB{
		Nombre:      periodo.GetNombre(),
		FechaInicio: periodo.GetFechaInicio(),
		FechaFin:    periodo.GetFechaFin(),
	}

	result := r.db.Table("periodos").Create(&periodoData)
	if result.Error != nil {
		return result.Error
	}

	periodo.SetID(periodoData.ID)

	return nil
}

func (r *PeriodoDao) Update(periodo *domain.Periodo) error {
	return r.db.Table("periodos").Where("id = ?", periodo.GetID()).Updates(periodoDB{
		Nombre:      periodo.GetNombre(),
		FechaInicio: periodo.GetFechaInicio(),
		FechaFin:    periodo.GetFechaFin(),
	}).Error
}

func (r *PeriodoDao) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&periodoDB{}).Error
}

func (r *PeriodoDao) AgregarCurso(periodoID, cursoID int64) error {
	query := "INSERT INTO periodo_cursos (periodo_id, curso_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE periodo_id = periodo_id"
	err := r.db.Exec(query, periodoID, cursoID).Error
	if err != nil {
		log.Printf("❌ Error al asociar curso %d al periodo %d: %v", cursoID, periodoID, err)
	}
	return err
}

func (r *PeriodoDao) ObtenerCursos(periodoID int64) ([]dto.CursoPeriodoDTO, error) {
	var cursos []dto.CursoPeriodoDTO

	query := `
		SELECT 
			pc.periodo_id, 
			pc.curso_id, 
			c.nombre, 
			c.edad_minima, 
			c.edad_maxima, 
			c.estado
		FROM periodo_cursos pc
		JOIN cursos c ON pc.curso_id = c.id
		WHERE pc.periodo_id = ?`

	result := r.db.Raw(query, periodoID).Scan(&cursos)
	if result.Error != nil {
		return nil, result.Error
	}

	return cursos, nil
}
