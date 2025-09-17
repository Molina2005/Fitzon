package crearcuenta

import (
	"database/sql"
	"log"
	"net/http"
)

func ProcesarCrearCuenta(db *sql.DB) http.HandlerFunc{
	return func (wr http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost{
			http.Error(wr, "Metodo no permitido", http.StatusMethodNotAllowed)
			return
		}

		nombre__usuario := r.FormValue("nombre")
		apellido__usuario := r.FormValue("apellido")
		correo__usuario := r.FormValue("email")
		contraseña__usuario := r.FormValue("password")

		_,err := db.Exec("INSERT INTO usuarios (nombre, apellido, correo_electronico, contrasena) VALUES ($1,$2,$3,$4)", nombre__usuario, apellido__usuario,correo__usuario, contraseña__usuario)
		if err != nil {
			log.Fatal(err)
		}
	}
}

