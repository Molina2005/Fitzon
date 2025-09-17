package connect

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConexionDB()(*sql.DB, error){
	// Cargue arhivo .env
	err := godotenv.Load()
	if err != nil{
		log.Fatal(errors.New("error al cargar archivo .env"))
	}

	// Cadena conexion DB 
	infBD := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Abrir conexion
	db, err := sql.Open( "postgres",infBD)
	if err != nil {
		return nil,err
	}

	// verificacion de la conexion
	if err := db.Ping();err != nil{
		return nil,err
	}
	return db, nil
}
