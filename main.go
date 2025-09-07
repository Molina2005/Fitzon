package main

import (
	"fmt"
	"log"

	"github.com/postgres/fitzon/connect"
)

func main(){
	conexion, err := conexiondb.ConexionDB()
	if err != nil {
		log.Fatal(err)
	}
	// cierre de conexion
	defer conexion.Close()
	fmt.Println("Conexion exitosa a PosgreSQL")
}