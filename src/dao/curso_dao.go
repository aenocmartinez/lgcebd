package dao

import (
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
	return r.db.Table("cursos").Create(curso).Error
}

func (r *CursoDao) Update(curso *domain.Curso) error {
	return r.db.Table("cursos").Where("id = ?", curso.GetID()).Updates(curso).Error
}

func (r *CursoDao) Delete(id int64) error {
	return r.db.Table("cursos").Where("id = ?", id).Delete(&domain.Curso{}).Error
}

func (r *CursoDao) FindByID(id int64) (*domain.Curso, error) {
	var cursoData cursoDB
	result := r.db.Table("cursos").Where("id = ?", id).First(&cursoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewCurso(nil), nil
		}
		return domain.NewCurso(nil), result.Error
	}

	return r.convertToDomain(&cursoData), nil
}

func (r *CursoDao) FindByNombre(nombre string) (*domain.Curso, error) {
	var cursoData cursoDB
	result := r.db.Table("cursos").Where("nombre = ?", nombre).First(&cursoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewCurso(nil), nil
		}
		return domain.NewCurso(nil), result.Error
	}

	return r.convertToDomain(&cursoData), nil
}

func (r *CursoDao) List() ([]dto.CursoDTO, error) {
	var cursos []dto.CursoDTO
	result := r.db.Table("cursos").Select("id, nombre, edad_minima, edad_maxima, estado").Find(&cursos)
	if result.Error != nil {
		return []dto.CursoDTO{}, result.Error
	}
	return cursos, nil
}

type cursoDB struct {
	ID         int64
	Nombre     string
	EdadMinima int
	EdadMaxima int
	Estado     string
}

func (r *CursoDao) convertToDomain(cursoData *cursoDB) *domain.Curso {
	curso := domain.NewCurso(nil)
	curso.SetID(cursoData.ID)
	curso.SetNombre(cursoData.Nombre)
	curso.SetEdadMinima(cursoData.EdadMinima)
	curso.SetEdadMaxima(cursoData.EdadMaxima)
	curso.SetEstado(cursoData.Estado)
	return curso
}
