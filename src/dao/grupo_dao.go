package dao

import (
	"ebd/src/domain"
	"log"

	"gorm.io/gorm"
)

type GrupoDao struct {
	db *gorm.DB
}

func NewGrupoDao(db *gorm.DB) *GrupoDao {
	return &GrupoDao{db: db}
}

type grupoInsertDB struct {
	ID             int64 `gorm:"column:id;primaryKey;autoIncrement"`
	CelebracionID  int64 `gorm:"column:celebracion_id"`
	PeriodoCursoID int64 `gorm:"column:periodo_curso_id"`
}

type grupoDB struct {
	ID                int64  `gorm:"column:id"`
	CelebracionID     int64  `gorm:"column:celebracion_id"`
	PeriodoCursoID    int64  `gorm:"column:periodo_curso_id"`
	NombreCelebracion string `gorm:"column:nombre_celebracion"`
	NombreCurso       string `gorm:"column:nombre_curso"`
	EdadMinima        int    `gorm:"column:edad_minima"`
	EdadMaxima        int    `gorm:"column:edad_maxima"`
	EstadoCurso       string `gorm:"column:estado_curso"`
	NombrePeriodo     string `gorm:"column:nombre_periodo"`
	IDPeriodo         int64  `gorm:"column:periodo_id"`
	FechaInicio       string `gorm:"column:fecha_inicio"`
	FechaFin          string `gorm:"column:fecha_fin"`
}

func (grupoDB) TableName() string {
	return "grupos"
}

func (grupoInsertDB) TableName() string {
	return "grupos"
}

func (r *GrupoDao) FindByID(id int64) *domain.Grupo {
	var grupoData grupoDB
	result := r.db.Where("id = ?", id).First(&grupoData)
	if result.Error != nil {
		return domain.NewGrupo(r)
	}

	grupo := domain.NewGrupo(r)
	grupo.SetID(grupoData.ID)
	grupo.SetCelebracion(domain.NewCelebracion(nil))
	grupo.GetCelebracion().SetID(grupoData.CelebracionID)
	grupo.SetCursoPeriodo(domain.NewCursoPeriodoEmpty())
	grupo.GetCursoPeriodo().SetID(grupoData.PeriodoCursoID)

	return grupo
}

func (r *GrupoDao) FindByCursoPeriodoYCelebracion(cursoPeriodoID, celebracionID int64) *domain.Grupo {
	var grupoData grupoDB
	result := r.db.Where("periodo_curso_id = ? AND celebracion_id = ?", cursoPeriodoID, celebracionID).First(&grupoData)
	if result.Error != nil {
		return domain.NewGrupo(r)
	}

	grupo := domain.NewGrupo(r)
	grupo.SetID(grupoData.ID)
	grupo.SetCelebracion(domain.NewCelebracion(nil))
	grupo.GetCelebracion().SetID(grupoData.CelebracionID)
	grupo.SetCursoPeriodo(domain.NewCursoPeriodoEmpty())
	grupo.GetCursoPeriodo().SetID(grupoData.PeriodoCursoID)

	return grupo
}

func (r *GrupoDao) Save(grupo *domain.Grupo) error {
	grupoData := grupoInsertDB{
		CelebracionID:  grupo.GetCelebracion().GetID(),
		PeriodoCursoID: grupo.GetCursoPeriodo().GetID(),
	}

	result := r.db.Create(&grupoData)
	if result.Error != nil {
		log.Println(result)
		return result.Error
	}

	grupo.SetID(grupoData.ID)

	return nil
}

func (r *GrupoDao) Update(grupo *domain.Grupo) error {
	return r.db.Where("id = ?", grupo.GetID()).Updates(grupoDB{
		CelebracionID:  grupo.GetCelebracion().GetID(),
		PeriodoCursoID: grupo.GetCursoPeriodo().GetID(),
	}).Error
}

func (r *GrupoDao) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&grupoDB{}).Error
}

func (r *GrupoDao) List() []domain.Grupo {
	var gruposData []grupoDB
	query := `
		SELECT g.id, g.celebracion_id, g.periodo_curso_id, 
			c.nombre AS nombre_curso, c.edad_minima, c.edad_maxima, c.estado AS estado_curso,
			p.id AS periodo_id, p.nombre AS nombre_periodo, p.fecha_inicio, p.fecha_fin,
			ce.nombre AS nombre_celebracion
		FROM grupos g
		JOIN periodo_cursos pc ON g.periodo_curso_id = pc.id
		JOIN cursos c ON pc.curso_id = c.id
		JOIN periodos p ON pc.periodo_id = p.id
		JOIN celebraciones ce ON g.celebracion_id = ce.id
		ORDER BY g.id ASC`

	result := r.db.Raw(query).Scan(&gruposData)
	if result.Error != nil {
		return nil
	}

	var grupos []domain.Grupo
	for _, grupo := range gruposData {
		g := domain.NewGrupo(r)
		g.SetID(grupo.ID)

		celebracion := domain.NewCelebracion(nil)
		celebracion.SetID(grupo.CelebracionID)
		celebracion.SetNombre(grupo.NombreCelebracion)
		g.SetCelebracion(celebracion)

		curso := domain.NewCurso(nil)
		curso.SetID(grupo.PeriodoCursoID)
		curso.SetNombre(grupo.NombreCurso)
		curso.SetEdadMinima(grupo.EdadMinima)
		curso.SetEdadMaxima(grupo.EdadMaxima)
		curso.SetEstado(grupo.EstadoCurso)

		periodo := domain.NewPeriodo(nil)
		periodo.SetID(grupo.IDPeriodo)
		periodo.SetNombre(grupo.NombrePeriodo)
		periodo.SetFechaInicio(grupo.FechaInicio)
		periodo.SetFechaFin(grupo.FechaFin)

		cursoPeriodo := domain.NewCursoPeriodo(grupo.PeriodoCursoID, curso, periodo)
		g.SetCursoPeriodo(cursoPeriodo)

		grupos = append(grupos, *g)
	}

	return grupos
}

func (r *GrupoDao) AgregarMaestro(grupoID int64, maestroID int64) error {
	query := "INSERT INTO grupo_maestros (grupo_id, maestro_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE grupo_id = grupo_id"
	return r.db.Exec(query, grupoID, maestroID).Error
}

func (r *GrupoDao) QuitarMaestros(grupoID int64) error {
	return r.db.Where("grupo_id = ?", grupoID).Delete(&domain.GrupoMaestro{}).Error
}

func (r *GrupoDao) ListarMaestros(grupoID int64) []domain.GrupoMaestro {
	var grupoMaestrosData []struct {
		ID              int64  `gorm:"column:id"`
		GrupoID         int64  `gorm:"column:grupo_id"`
		MaestroID       int64  `gorm:"column:maestro_id"`
		Nombre          string `gorm:"column:nombre"`
		Telefono        string `gorm:"column:telefono"`
		FechaNacimiento string `gorm:"column:fecha_nacimiento"`
		Estado          string `gorm:"column:estado"`
	}

	query := `
		SELECT gm.id, gm.grupo_id, m.id AS maestro_id, 
			m.nombre, m.telefono, m.fecha_nacimiento, m.estado
		FROM grupo_maestros gm
		JOIN maestros m ON gm.maestro_id = m.id
		WHERE gm.grupo_id = ?`

	r.db.Raw(query, grupoID).Scan(&grupoMaestrosData)

	var grupoMaestros []domain.GrupoMaestro
	for _, gmData := range grupoMaestrosData {
		gm := domain.NewGrupoMaestro(r)
		gm.SetID(gmData.ID)
		gm.SetGrupo(domain.NewGrupo(nil))
		gm.GetGrupo().SetID(gmData.GrupoID)

		maestro := domain.NewMaestro(nil)
		maestro.SetID(gmData.MaestroID)
		maestro.SetNombre(gmData.Nombre)
		maestro.SetTelefono(gmData.Telefono)
		maestro.SetFechaNacimiento(gmData.FechaNacimiento)
		maestro.SetEstado(gmData.Estado)
		gm.SetMaestro(maestro)

		grupoMaestros = append(grupoMaestros, *gm)
	}

	return grupoMaestros
}
