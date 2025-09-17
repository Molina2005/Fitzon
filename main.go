package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/postgres/fitzon/connect"
	crearcuenta "github.com/postgres/fitzon/crear_cuenta"
)

func main(){
	// Conexion base de datos
	conexion, err := connect.ConexionDB()
	if err != nil {
		log.Fatal(err)
	}
	// cierre de conexion
	defer conexion.Close()
	fmt.Println("Conexion exitosa a PosgreSQL")

	// Formulario
	http.HandleFunc("/procesar", crearcuenta.ProcesarCrearCuenta(conexion))
	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))	
}