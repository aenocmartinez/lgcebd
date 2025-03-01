package dao

import (
	"ebd/src/domain"
	"fmt"

	"gorm.io/gorm"
)

type ClaseDao struct {
	db *gorm.DB
}

func NewClaseDao(db *gorm.DB) *ClaseDao {
	return &ClaseDao{db: db}
}

type claseDB struct {
	ID                  int64   `gorm:"column:id"`
	Fecha               string  `gorm:"column:fecha"`
	Ofrenda             float64 `gorm:"column:ofrenda"`
	GrupoID             int64   `gorm:"column:grupo_id"`
	ContenidoTematicoID int64   `gorm:"column:contenido_tematico_id"`
	CelebracionID       int64   `gorm:"column:celebracion_id"`
	NombreCelebracion   string  `gorm:"column:nombre_celebracion"`
	PeriodoCursoID      int64   `gorm:"column:periodo_curso_id"`
}

type insertClaseDB struct {
	ID                  int64   `gorm:"column:id"`
	Fecha               string  `gorm:"column:fecha"`
	Ofrenda             float64 `gorm:"column:ofrenda"`
	GrupoID             int64   `gorm:"column:grupo_id"`
	ContenidoTematicoID int64   `gorm:"column:contenido_tematico_id"`
}

func (c insertClaseDB) TableName() string {
	return "clases"
}

func (c claseDB) TableName() string {
	return "clases"
}

func (r *ClaseDao) FindByID(id int64) *domain.Clase {
	var claseData claseDB
	query := `
		SELECT c.id, c.fecha, c.ofrenda, c.grupo_id, c.contenido_tematico_id, 
		g.celebracion_id, ce.nombre AS nombre_celebracion, g.periodo_curso_id
		FROM clases c
		JOIN grupos g ON c.grupo_id = g.id
		JOIN celebraciones ce ON g.celebracion_id = ce.id
		WHERE c.id = ?`
	result := r.db.Raw(query, id).Scan(&claseData)
	if result.Error != nil {
		return domain.NewClase(r)
	}

	clase := domain.NewClase(r)
	clase.SetID(claseData.ID)
	clase.SetFecha(claseData.Fecha)
	clase.SetOfrenda(claseData.Ofrenda)

	grupo := domain.NewGrupo(nil)
	grupo.SetID(claseData.GrupoID)

	celebracion := domain.NewCelebracion(nil)
	celebracion.SetID(claseData.CelebracionID)
	celebracion.SetNombre(claseData.NombreCelebracion)
	grupo.SetCelebracion(celebracion)

	cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
	cursoPeriodo.SetID(claseData.PeriodoCursoID)
	grupo.SetCursoPeriodo(cursoPeriodo)

	clase.SetGrupo(grupo)

	contenido := domain.NewContenidoTematico(nil)
	contenido.SetID(claseData.ContenidoTematicoID)
	clase.SetContenidoTematico(contenido)

	return clase
}

func (r *ClaseDao) FindByGrupoFecha(grupoID int64, fecha string) *domain.Clase {
	var claseData claseDB
	query := `
		SELECT c.id, c.fecha, c.ofrenda, c.grupo_id, c.contenido_tematico_id, 
		g.celebracion_id, ce.nombre AS nombre_celebracion, g.periodo_curso_id
		FROM clases c
		JOIN grupos g ON c.grupo_id = g.id
		JOIN celebraciones ce ON g.celebracion_id = ce.id
		WHERE c.grupo_id = ? AND c.fecha = ?`
	result := r.db.Raw(query, grupoID, fecha).Scan(&claseData)
	if result.Error != nil {
		return domain.NewClase(r)
	}

	clase := domain.NewClase(r)
	clase.SetID(claseData.ID)
	clase.SetFecha(claseData.Fecha)
	clase.SetOfrenda(claseData.Ofrenda)

	grupo := domain.NewGrupo(nil)
	grupo.SetID(claseData.GrupoID)

	celebracion := domain.NewCelebracion(nil)
	celebracion.SetID(claseData.CelebracionID)
	celebracion.SetNombre(claseData.NombreCelebracion)
	grupo.SetCelebracion(celebracion)

	cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
	cursoPeriodo.SetID(claseData.PeriodoCursoID)
	grupo.SetCursoPeriodo(cursoPeriodo)

	clase.SetGrupo(grupo)

	contenido := domain.NewContenidoTematico(nil)
	contenido.SetID(claseData.ContenidoTematicoID)
	clase.SetContenidoTematico(contenido)

	return clase
}

func (r *ClaseDao) List() ([]domain.Clase, error) {
	var clasesData []claseDB
	query := `
		SELECT c.id, c.fecha, c.ofrenda, c.grupo_id, c.contenido_tematico_id, 
		g.celebracion_id, ce.nombre AS nombre_celebracion, g.periodo_curso_id
		FROM clases c
		JOIN grupos g ON c.grupo_id = g.id
		JOIN celebraciones ce ON g.celebracion_id = ce.id`
	result := r.db.Raw(query).Scan(&clasesData)
	if result.Error != nil {
		return nil, result.Error
	}

	var clases []domain.Clase
	for _, claseData := range clasesData {
		clase := domain.NewClase(r)
		clase.SetID(claseData.ID)
		clase.SetFecha(claseData.Fecha)
		clase.SetOfrenda(claseData.Ofrenda)

		grupo := domain.NewGrupo(nil)
		grupo.SetID(claseData.GrupoID)

		celebracion := domain.NewCelebracion(nil)
		celebracion.SetID(claseData.CelebracionID)
		celebracion.SetNombre(claseData.NombreCelebracion)
		grupo.SetCelebracion(celebracion)

		cursoPeriodo := domain.NewCursoPeriodoEmpty(nil)
		cursoPeriodo.SetID(claseData.PeriodoCursoID)
		grupo.SetCursoPeriodo(cursoPeriodo)

		clase.SetGrupo(grupo)

		contenido := domain.NewContenidoTematico(nil)
		contenido.SetID(claseData.ContenidoTematicoID)
		clase.SetContenidoTematico(contenido)

		clases = append(clases, *clase)
	}

	return clases, nil
}

func (r *ClaseDao) Save(clase *domain.Clase) error {
	claseData := insertClaseDB{
		Fecha:               clase.GetFecha(),
		Ofrenda:             clase.GetOfrenda(),
		GrupoID:             clase.GetGrupo().GetID(),
		ContenidoTematicoID: clase.GetContenidoTematico().GetID(),
	}

	result := r.db.Table("clases").Create(&claseData)
	if result.Error != nil {
		fmt.Println(result)
		return result.Error
	}

	clase.SetID(claseData.ID)
	return nil
}

func (r *ClaseDao) Update(clase *domain.Clase) error {
	return r.db.Table("clases").Where("id = ?", clase.GetID()).Updates(map[string]interface{}{
		"fecha":                 clase.GetFecha(),
		"ofrenda":               clase.GetOfrenda(),
		"grupo_id":              clase.GetGrupo().GetID(),
		"contenido_tematico_id": clase.GetContenidoTematico().GetID(),
	}).Error
}

func (r *ClaseDao) Delete(claseID int64) error {
	return r.db.Table("clases").Where("id = ?", claseID).Delete(nil).Error
}

func (r *ClaseDao) RegistrarAsistencia(claseID int64, matriculaID int64) error {
	query := "INSERT INTO asistencias (clase_id, matricula_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE clase_id = clase_id"
	return r.db.Exec(query, claseID, matriculaID).Error
}
