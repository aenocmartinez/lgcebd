package dao

import (
	"database/sql"
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
	ID             int64         `gorm:"column:id"`
	Descripcion    string        `gorm:"column:descripcion"`
	PeriodoCursoID int64         `gorm:"column:periodo_curso_id"`
	Orden          sql.NullInt64 `gorm:"column:orden"`
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
	if contenidoData.Orden.Valid {
		contenido.SetOrden(int(contenidoData.Orden.Int64))
	} else {
		contenido.SetOrden(0)
	}

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
	if contenidoData.Orden.Valid {
		contenido.SetOrden(int(contenidoData.Orden.Int64))
	} else {
		contenido.SetOrden(0)
	}

	return contenido
}

func (r *ContenidoTematicoDao) Save(contenido *domain.ContenidoTematico) error {
	ordenValue := sql.NullInt64{Valid: false}
	if contenido.GetOrden() != 0 {
		ordenValue = sql.NullInt64{Int64: int64(contenido.GetOrden()), Valid: true}
	}

	contenidoData := contenidoTematicoDB{
		Descripcion:    contenido.GetDescripcion(),
		PeriodoCursoID: contenido.GetCursoPeriodo().GetID(),
		Orden:          ordenValue,
	}

	result := r.db.Create(&contenidoData)
	if result.Error != nil {
		return result.Error
	}

	contenido.SetID(contenidoData.ID)
	return nil
}

func (r *ContenidoTematicoDao) Update(contenido *domain.ContenidoTematico) error {
	ordenValue := sql.NullInt64{Valid: false}
	if contenido.GetOrden() != 0 {
		ordenValue = sql.NullInt64{Int64: int64(contenido.GetOrden()), Valid: true}
	}

	return r.db.Where("id = ?", contenido.GetID()).Updates(contenidoTematicoDB{
		Descripcion:    contenido.GetDescripcion(),
		PeriodoCursoID: contenido.GetCursoPeriodo().GetID(),
		Orden:          ordenValue,
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
		if contenido.Orden.Valid {
			c.SetOrden(int(contenido.Orden.Int64))
		} else {
			c.SetOrden(0)
		}
		contenidos = append(contenidos, *c)
	}

	return contenidos
}
