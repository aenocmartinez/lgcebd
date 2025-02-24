package di

import (
	"sync"

	"ebd/src/dao"
	"ebd/src/infraestructure/database"

	"gorm.io/gorm"
)

type Container struct {
	db               *gorm.DB
	userRepo         *dao.UserDao
	cursoRepo        *dao.CursoDao
	periodoRepo      *dao.PeriodoDao
	alumnoRepo       *dao.AlumnoDao
	cursoPeriodoRepo *dao.CursoPeriodoDao
	matriculaRepo    *dao.MatriculaDao
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:               db,
			userRepo:         dao.NewUserDao(db),
			cursoRepo:        dao.NewCursoDao(db),
			periodoRepo:      dao.NewPeriodoDao(db),
			alumnoRepo:       dao.NewAlumnoDao(db),
			cursoPeriodoRepo: dao.NewCursoPeriodoDao(db),
			matriculaRepo:    dao.NewMatriculaDao(db),
		}
	})
	return instance
}

func (c *Container) GetUserRepository() *dao.UserDao {
	return c.userRepo
}

func (c *Container) GetCursoRepository() *dao.CursoDao {
	return c.cursoRepo
}

func (c *Container) GetPeriodoRepository() *dao.PeriodoDao {
	return c.periodoRepo
}

func (c *Container) GetAlumnoRepository() *dao.AlumnoDao {
	return c.alumnoRepo
}

func (c *Container) GetCursoPeriodoRepository() *dao.CursoPeriodoDao {
	return c.cursoPeriodoRepo
}

func (c *Container) GetMatriculaRepository() *dao.MatriculaDao {
	return c.matriculaRepo
}
