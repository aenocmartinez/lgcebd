package dao

import (
	"ebd/src/domain"
	"fmt"

	"gorm.io/gorm"
)

type MatriculaDao struct {
	db *gorm.DB
}

func NewMatriculaDao(db *gorm.DB) *MatriculaDao {
	return &MatriculaDao{db: db}
}

type matriculaDB struct {
	ID             int64 `gorm:"column:id"`
	AlumnoID       int64 `gorm:"column:alumnno_id"`
	CursoPeriodoID int64 `gorm:"column:periodo_curso_id"`
}

func (r *MatriculaDao) FindByID(id int64) (*domain.Matricula, error) {
	var matriculaData matriculaDB
	result := r.db.Table("matriculas").Where("id = ?", id).First(&matriculaData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.NewMatriculaEmpty(), nil
		}
		return nil, result.Error
	}

	alumno := domain.NewAlumno(nil)
	alumno.SetID(matriculaData.AlumnoID)

	cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
	cursoPeriodo.SetID(matriculaData.CursoPeriodoID)

	return domain.NewMatricula(matriculaData.ID, alumno, cursoPeriodo), nil
}

func (r *MatriculaDao) ExisteMatricula(alumnoID, cursoPeriodoID int64) bool {
	var count int64
	r.db.Table("matriculas").
		Where("alumnno_id = ? AND periodo_curso_id = ?", alumnoID, cursoPeriodoID).
		Count(&count)
	return count > 0
}

func (r *MatriculaDao) Save(matricula *domain.Matricula) error {
	matriculaData := matriculaDB{
		AlumnoID:       matricula.GetAlumnoID(),
		CursoPeriodoID: matricula.GetCursoPeriodoID(),
	}

	result := r.db.Table("matriculas").Create(&matriculaData)
	if result.Error != nil {
		fmt.Println(result)
		return result.Error
	}

	matricula.SetID(matriculaData.ID)
	return nil
}

func (r *MatriculaDao) ObtenerMatriculasPorAlumno(alumnoID int64) ([]domain.Matricula, error) {
	var matriculasData []matriculaDB
	result := r.db.Table("matriculas").Where("alumnno_id = ?", alumnoID).Find(&matriculasData)
	if result.Error != nil {
		return nil, result.Error
	}

	var matriculas []domain.Matricula
	for _, m := range matriculasData {
		alumno := domain.NewAlumno(nil)
		alumno.SetID(m.AlumnoID)

		cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
		cursoPeriodo.SetID(m.CursoPeriodoID)

		matriculas = append(matriculas, *domain.NewMatricula(m.ID, alumno, cursoPeriodo))
	}

	return matriculas, nil
}

func (r *MatriculaDao) Delete(id int64) error {
	return r.db.Table("matriculas").Where("id = ?", id).Delete(&matriculaDB{}).Error
}

func (r *MatriculaDao) ObtenerMatriculasPorCursoPeriodo(cursoPeriodoID int64) ([]domain.Matricula, error) {
	var matriculasData []struct {
		ID                int64  `gorm:"column:id"`
		AlumnoID          int64  `gorm:"column:alumnno_id"`
		Nombre            string `gorm:"column:nombre"`
		FechaNacimiento   string `gorm:"column:fecha_nacimiento"`
		Telefono          string `gorm:"column:telefono"`
		Acudiente         string `gorm:"column:acudiente"`
		AcudienteTelefono string `gorm:"column:acudiente_telefono"`
		Direccion         string `gorm:"column:direccion"`
		CursoPeriodoID    int64  `gorm:"column:periodo_curso_id"`
	}

	query := `
		SELECT m.id, 
		       a.id AS alumnno_id, 
		       a.nombre, 
		       a.fecha_nacimiento, 
		       a.telefono, 
		       a.acudiente, 
		       a.acudiente_telefono, 
		       a.direccion,
		       m.periodo_curso_id
		FROM matriculas m
		INNER JOIN alumnos a ON a.id = m.alumnno_id
		WHERE m.periodo_curso_id = ?
	`

	result := r.db.Raw(query, cursoPeriodoID).Scan(&matriculasData)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var matriculas []domain.Matricula
	for _, data := range matriculasData {
		alumno := domain.NewAlumno(nil)
		alumno.SetID(data.AlumnoID)
		alumno.SetNombre(data.Nombre)
		alumno.SetFechaNacimiento(data.FechaNacimiento)
		alumno.SetTelefono(data.Telefono)
		alumno.SetAcudiente(data.Acudiente)
		alumno.SetAcudienteTelefono(data.AcudienteTelefono)
		alumno.SetDireccion(data.Direccion)

		cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
		cursoPeriodo.SetID(data.CursoPeriodoID)

		matricula := domain.NewMatricula(data.ID, alumno, cursoPeriodo)
		matriculas = append(matriculas, *matricula)
	}

	return matriculas, nil
}
