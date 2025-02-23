package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {

		// if err := godotenv.Load(); err != nil {
		// 	log.Fatalf("Error cargando el archivo .env: %v", err)
		// }

		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
			"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
			os.Getenv("DB_NAME") + "?parseTime=true"

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error al conectar con la base de datos: %v", err)
		}

		log.Println("Conexión a la base de datos establecida con GORM.")
	})

	return db
}
