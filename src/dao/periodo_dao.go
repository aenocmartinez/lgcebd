package dao

import (
	"ebd/src/domain"
	"ebd/src/view/dto"

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
