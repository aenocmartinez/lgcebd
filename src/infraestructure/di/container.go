package di

import (
	"sync"

	"ebd/src/dao"
	"ebd/src/infraestructure/database"

	"gorm.io/gorm"
)

type Container struct {
	db        *gorm.DB
	userRepo  *dao.UserDao
	cursoRepo *dao.CursoDao
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:        db,
			userRepo:  dao.NewUserDao(db),
			cursoRepo: dao.NewCursoDao(db), // ✅ Se inyecta CursoDao
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
