package dao

import (
	"errors"

	"ebd/src/domain"
	"ebd/src/view/dto"

	"gorm.io/gorm"
)

type CursoDao struct {
	db *gorm.DB
}

func NewCursoDao(db *gorm.DB) *CursoDao {
	return &CursoDao{db: db}
}

func (r *CursoDao) Save(curso *domain.Curso) error {
	cursoData := cursoDB{
		Nombre:     curso.GetNombre(),
		EdadMinima: curso.GetEdadMinima(),
		EdadMaxima: curso.GetEdadMaxima(),
		Estado:     curso.GetEstado(),
	}

	err := r.db.Table("cursos").Create(&cursoData).Error
	if err != nil {
		return err
	}

	curso.SetID(cursoData.ID)

	return nil
}

func (r *CursoDao) Update(curso *domain.Curso) error {
	cursoData := cursoDB{
		Nombre:     curso.GetNombre(),
		EdadMinima: curso.GetEdadMinima(),
		EdadMaxima: curso.GetEdadMaxima(),
		Estado:     curso.GetEstado(),
	}

	return r.db.Table("cursos").Where("id = ?", curso.GetID()).Updates(&cursoData).Error
}

func (r *CursoDao) FindByID(id int64) (*domain.Curso, error) {
	var cursoData cursoDB
	result := r.db.Table("cursos").Where("id = ?", id).First(&cursoData)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.NewCurso(r), nil
		}
		return nil, result.Error
	}

	curso := domain.NewCurso(r)
	curso.SetID(cursoData.ID)
	curso.SetNombre(cursoData.Nombre)
	curso.SetEdadMinima(cursoData.EdadMinima)
	curso.SetEdadMaxima(cursoData.EdadMaxima)
	curso.SetEstado(cursoData.Estado)

	return curso, nil
}

func (r *CursoDao) FindByNombre(nombre string) (*domain.Curso, error) {
	var cursoData cursoDB
	result := r.db.Table("cursos").Where("nombre = ?", nombre).First(&cursoData)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.NewCurso(r), nil
		}
		return nil, result.Error
	}

	curso := domain.NewCurso(r)
	curso.SetID(cursoData.ID)
	curso.SetNombre(cursoData.Nombre)
	curso.SetEdadMinima(cursoData.EdadMinima)
	curso.SetEdadMaxima(cursoData.EdadMaxima)
	curso.SetEstado(cursoData.Estado)

	return curso, nil
}

func (r *CursoDao) Delete(id int64) error {
	return r.db.Table("cursos").Where("id = ?", id).Delete(&cursoDB{}).Error
}

func (r *CursoDao) List() ([]dto.CursoDTO, error) {
	var cursosData []cursoDB
	result := r.db.Table("cursos").Find(&cursosData)

	if result.Error != nil {
		return []dto.CursoDTO{}, result.Error
	}

	cursosDTO := []dto.CursoDTO{}
	for _, cursoData := range cursosData {
		curso := domain.NewCurso(r)
		curso.SetID(cursoData.ID)
		curso.SetNombre(cursoData.Nombre)
		curso.SetEdadMinima(cursoData.EdadMinima)
		curso.SetEdadMaxima(cursoData.EdadMaxima)
		curso.SetEstado(cursoData.Estado)

		cursosDTO = append(cursosDTO, *curso.ToDTO())
	}

	return cursosDTO, nil
}

type cursoDB struct {
	ID         int64  `gorm:"column:id"`
	Nombre     string `gorm:"column:nombre"`
	EdadMinima int    `gorm:"column:edad_minima"`
	EdadMaxima int    `gorm:"column:edad_maxima"`
	Estado     string `gorm:"column:estado"`
}
