package dao

import (
	"ebd/src/domain"
	"ebd/src/view/dto"

	"gorm.io/gorm"
)

type AlumnoDao struct {
	db *gorm.DB
}

func NewAlumnoDao(db *gorm.DB) *AlumnoDao {
	return &AlumnoDao{db: db}
}

type alumnoDB struct {
	ID                int64  `gorm:"column:id"`
	Nombre            string `gorm:"column:nombre"`
	FechaNacimiento   string `gorm:"column:fecha_nacimiento"`
	Telefono          string `gorm:"column:telefono"`
	Acudiente         string `gorm:"column:acudiente"`
	AcudienteTelefono string `gorm:"column:acudiente_telefono"`
	Direccion         string `gorm:"column:direccion"`
}

func (alumnoDB) TableName() string {
	return "alumnos"
}

func (r *AlumnoDao) FindByID(id int64) (*domain.Alumno, error) {
	var alumnoData alumnoDB
	result := r.db.Where("id = ?", id).First(&alumnoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &domain.Alumno{}, nil
		}
		return nil, result.Error
	}

	alumno := domain.NewAlumno(r)
	alumno.SetID(alumnoData.ID)
	alumno.SetNombre(alumnoData.Nombre)
	alumno.SetFechaNacimiento(alumnoData.FechaNacimiento)
	alumno.SetTelefono(alumnoData.Telefono)
	alumno.SetAcudiente(alumnoData.Acudiente)
	alumno.SetAcudienteTelefono(alumnoData.AcudienteTelefono)
	alumno.SetDireccion(alumnoData.Direccion)

	return alumno, nil
}

func (r *AlumnoDao) FindByNombre(nombre string) (*domain.Alumno, error) {
	var alumnoData alumnoDB
	result := r.db.Where("nombre = ?", nombre).First(&alumnoData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return &domain.Alumno{}, nil
		}
		return nil, result.Error
	}

	alumno := domain.NewAlumno(r)
	alumno.SetID(alumnoData.ID)
	alumno.SetNombre(alumnoData.Nombre)
	alumno.SetFechaNacimiento(alumnoData.FechaNacimiento)
	alumno.SetTelefono(alumnoData.Telefono)
	alumno.SetAcudiente(alumnoData.Acudiente)
	alumno.SetAcudienteTelefono(alumnoData.AcudienteTelefono)
	alumno.SetDireccion(alumnoData.Direccion)

	return alumno, nil
}

func (r *AlumnoDao) List() ([]dto.AlumnoDTO, error) {
	var alumnosData []alumnoDB
	result := r.db.Find(&alumnosData)
	if result.Error != nil {
		return nil, result.Error
	}

	alumnosDTO := []dto.AlumnoDTO{}
	for _, alumno := range alumnosData {
		alumnosDTO = append(alumnosDTO, dto.AlumnoDTO{
			ID:                alumno.ID,
			Nombre:            alumno.Nombre,
			FechaNacimiento:   alumno.FechaNacimiento,
			Telefono:          alumno.Telefono,
			Acudiente:         alumno.Acudiente,
			AcudienteTelefono: alumno.AcudienteTelefono,
			Direccion:         alumno.Direccion,
		})
	}

	return alumnosDTO, nil
}

func (r *AlumnoDao) Save(alumno *domain.Alumno) error {
	alumnoData := alumnoDB{
		Nombre:            alumno.GetNombre(),
		FechaNacimiento:   alumno.GetFechaNacimiento(),
		Telefono:          alumno.GetTelefono(),
		Acudiente:         alumno.GetAcudiente(),
		AcudienteTelefono: alumno.GetAcudienteTelefono(),
		Direccion:         alumno.GetDireccion(),
	}

	result := r.db.Create(&alumnoData)
	if result.Error != nil {
		return result.Error
	}

	alumno.SetID(alumnoData.ID)
	return nil
}

func (r *AlumnoDao) Update(alumno *domain.Alumno) error {
	return r.db.Where("id = ?", alumno.GetID()).Updates(alumnoDB{
		Nombre:            alumno.GetNombre(),
		FechaNacimiento:   alumno.GetFechaNacimiento(),
		Telefono:          alumno.GetTelefono(),
		Acudiente:         alumno.GetAcudiente(),
		AcudienteTelefono: alumno.GetAcudienteTelefono(),
		Direccion:         alumno.GetDireccion(),
	}).Error
}

func (r *AlumnoDao) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&alumnoDB{}).Error
}
