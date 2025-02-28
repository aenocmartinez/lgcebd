package dao

import (
	"ebd/src/domain"

	"gorm.io/gorm"
)

type MaestroDao struct {
	db *gorm.DB
}

func NewMaestroDao(db *gorm.DB) *MaestroDao {
	return &MaestroDao{db: db}
}

type maestroDB struct {
	ID              int64  `gorm:"column:id"`
	Nombre          string `gorm:"column:nombre"`
	Telefono        string `gorm:"column:telefono"`
	FechaNacimiento string `gorm:"column:fecha_nacimiento"`
	Estado          string `gorm:"column:estado"`
}

func (r *MaestroDao) FindByID(id int64) *domain.Maestro {
	var maestroData maestroDB
	result := r.db.Table("maestros").Where("id = ?", id).First(&maestroData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewMaestro(r)
		}
		return domain.NewMaestro(r)
	}

	maestro := domain.NewMaestro(r)
	maestro.SetID(maestroData.ID)
	maestro.SetNombre(maestroData.Nombre)
	maestro.SetTelefono(maestroData.Telefono)
	maestro.SetFechaNacimiento(maestroData.FechaNacimiento)
	maestro.SetEstado(maestroData.Estado)

	return maestro
}

func (r *MaestroDao) List() ([]domain.Maestro, error) {
	var maestrosData []maestroDB
	result := r.db.Table("maestros").Find(&maestrosData)
	if result.Error != nil {
		return nil, result.Error
	}

	var maestros []domain.Maestro
	for _, data := range maestrosData {
		maestro := domain.NewMaestro(r) // Se inicializa con el repositorio
		maestro.SetID(data.ID)
		maestro.SetNombre(data.Nombre)
		maestro.SetTelefono(data.Telefono)
		maestro.SetFechaNacimiento(data.FechaNacimiento)
		maestro.SetEstado(data.Estado)
		maestros = append(maestros, *maestro)
	}

	return maestros, nil
}

func (r *MaestroDao) Save(maestro *domain.Maestro) error {
	maestroData := maestroDB{
		Nombre:          maestro.GetNombre(),
		Telefono:        maestro.GetTelefono(),
		FechaNacimiento: maestro.GetFechaNacimiento(),
		Estado:          maestro.GetEstado(),
	}

	result := r.db.Table("maestros").Create(&maestroData)
	if result.Error != nil {
		return result.Error
	}

	maestro.SetID(maestroData.ID)

	return nil
}

func (r *MaestroDao) Update(maestro *domain.Maestro) error {
	return r.db.Table("maestros").Where("id = ?", maestro.GetID()).Updates(maestroDB{
		Nombre:          maestro.GetNombre(),
		Telefono:        maestro.GetTelefono(),
		FechaNacimiento: maestro.GetFechaNacimiento(),
		Estado:          maestro.GetEstado(), // <-- Se agrega el estado
	}).Error
}

func (r *MaestroDao) Delete(id int64) error {
	return r.db.Table("maestros").Where("id = ?", id).Delete(&maestroDB{}).Error
}
