package dao

import (
	"ebd/src/domain"

	"gorm.io/gorm"
)

type ContenidoTematicoDao struct {
	db *gorm.DB
}

func NewContenidoTematicoDao(db *gorm.DB) *ContenidoTematicoDao {
	return &ContenidoTematicoDao{db: db}
}

type contenidoTematicoDB struct {
	ID             int64  `gorm:"column:id"`
	Descripcion    string `gorm:"column:descripcion"`
	PeriodoCursoID int64  `gorm:"column:periodo_curso_id"`
}

func (contenidoTematicoDB) TableName() string {
	return "contenido_tematico"
}

func (r *ContenidoTematicoDao) FindByID(id int64) *domain.ContenidoTematico {
	var contenidoData contenidoTematicoDB
	result := r.db.Where("id = ?", id).First(&contenidoData)
	if result.Error != nil {
		return domain.NewContenidoTematico(r)
	}

	contenido := domain.NewContenidoTematico(r)
	contenido.SetID(contenidoData.ID)
	contenido.SetDescripcion(contenidoData.Descripcion)
	contenido.SetCursoPeriodo(domain.NewCursoPeriodoEmpty(nil))
	contenido.GetCursoPeriodo().SetID(contenidoData.PeriodoCursoID)

	return contenido
}

func (r *ContenidoTematicoDao) FindByDescripcion(cursoPeriodoID int64, descripcion string) *domain.ContenidoTematico {
	var contenidoData contenidoTematicoDB
	result := r.db.Where("periodo_curso_id = ? AND descripcion = ?", cursoPeriodoID, descripcion).First(&contenidoData)
	if result.Error != nil {
		return domain.NewContenidoTematico(r)
	}

	contenido := domain.NewContenidoTematico(r)
	contenido.SetID(contenidoData.ID)
	contenido.SetDescripcion(contenidoData.Descripcion)
	contenido.SetCursoPeriodo(domain.NewCursoPeriodoEmpty(nil))
	contenido.GetCursoPeriodo().SetID(contenidoData.PeriodoCursoID)

	return contenido
}

func (r *ContenidoTematicoDao) Save(contenido *domain.ContenidoTematico) error {
	contenidoData := contenidoTematicoDB{
		Descripcion:    contenido.GetDescripcion(),
		PeriodoCursoID: contenido.GetCursoPeriodo().GetID(),
	}

	result := r.db.Create(&contenidoData)
	if result.Error != nil {
		return result.Error
	}

	contenido.SetID(contenidoData.ID)
	return nil
}

func (r *ContenidoTematicoDao) Update(contenido *domain.ContenidoTematico) error {
	return r.db.Where("id = ?", contenido.GetID()).Updates(contenidoTematicoDB{
		Descripcion:    contenido.GetDescripcion(),
		PeriodoCursoID: contenido.GetCursoPeriodo().GetID(),
	}).Error
}

func (r *ContenidoTematicoDao) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&contenidoTematicoDB{}).Error
}

func (r *ContenidoTematicoDao) List() []domain.ContenidoTematico {
	var contenidosData []contenidoTematicoDB
	result := r.db.Order("id ASC").Find(&contenidosData)
	if result.Error != nil {
		return nil
	}

	var contenidos []domain.ContenidoTematico
	for _, contenido := range contenidosData {
		c := domain.NewContenidoTematico(r)
		c.SetID(contenido.ID)
		c.SetDescripcion(contenido.Descripcion)
		c.SetCursoPeriodo(domain.NewCursoPeriodoEmpty(nil))
		c.GetCursoPeriodo().SetID(contenido.PeriodoCursoID)
		contenidos = append(contenidos, *c)
	}

	return contenidos
}
