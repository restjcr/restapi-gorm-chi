package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=postgres password=1234 dbname=restapi-gorm-chi port=5432"
var dsn_tests = "host=localhost user=postgres password=1234 dbname=restapi-gorm-chi-tests port=5432"

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	log.Println("Conexion existosa a la base de datos")

	return db
}

func NewConnectionForTest() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn_tests), &gorm.Config{})

	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	log.Println("Conexion existosa a la base de datos")

	return db
}
