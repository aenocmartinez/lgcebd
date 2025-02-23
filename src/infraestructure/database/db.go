package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		// if err := godotenv.Load(); err != nil {
		// 	log.Fatalf("Error cargando el archivo .env: %v", err)
		// }

		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser, dbPassword, dbHost, dbPort, dbName)

		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error al conectar con la base de datos: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Error verificando la conexión a la base de datos: %v", err)
		}

		log.Println("Conexión a la base de datos establecida exitosamente.")
	})

	return db
}
