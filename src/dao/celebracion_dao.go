package dao

import (
	"ebd/src/domain"

	"gorm.io/gorm"
)

type CelebracionDao struct {
	db *gorm.DB
}

func NewCelebracionDao(db *gorm.DB) *CelebracionDao {
	return &CelebracionDao{db: db}
}

type celebracionDB struct {
	ID     int64  `gorm:"column:id"`
	Nombre string `gorm:"column:nombre"`
}

func (celebracionDB) TableName() string {
	return "celebraciones"
}

func (r *CelebracionDao) FindByID(id int64) *domain.Celebracion {
	var celebracionData celebracionDB
	result := r.db.Where("id = ?", id).First(&celebracionData)
	if result.Error != nil {
		return &domain.Celebracion{}
	}

	celebracion := domain.NewCelebracion(r)
	celebracion.SetID(celebracionData.ID)
	celebracion.SetNombre(celebracionData.Nombre)

	return celebracion
}

func (r *CelebracionDao) FindByNombre(nombre string) *domain.Celebracion {
	var celebracionData celebracionDB
	result := r.db.Where("nombre = ?", nombre).First(&celebracionData)
	if result.Error != nil {
		return &domain.Celebracion{}
	}

	celebracion := domain.NewCelebracion(r)
	celebracion.SetID(celebracionData.ID)
	celebracion.SetNombre(celebracionData.Nombre)

	return celebracion
}

func (r *CelebracionDao) Save(celebracion *domain.Celebracion) error {
	celebracionData := celebracionDB{
		Nombre: celebracion.GetNombre(),
	}

	result := r.db.Create(&celebracionData)
	if result.Error != nil {
		return result.Error
	}

	celebracion.SetID(celebracionData.ID)
	return nil
}

func (r *CelebracionDao) Update(celebracion *domain.Celebracion) error {
	return r.db.Where("id = ?", celebracion.GetID()).Updates(celebracionDB{
		Nombre: celebracion.GetNombre(),
	}).Error
}

func (r *CelebracionDao) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&celebracionDB{}).Error
}

func (r *CelebracionDao) List() []domain.Celebracion {
	var celebracionesData []celebracionDB
	result := r.db.Order("id ASC").Find(&celebracionesData)
	if result.Error != nil {
		return nil
	}

	var celebraciones []domain.Celebracion
	for _, celebracion := range celebracionesData {
		c := domain.NewCelebracion(r)
		c.SetID(celebracion.ID)
		c.SetNombre(celebracion.Nombre)
		celebraciones = append(celebraciones, *c)
	}

	return celebraciones
}
