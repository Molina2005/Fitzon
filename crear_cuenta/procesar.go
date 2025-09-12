package crearcuenta

import (
	"net/http"
)

// http.ResponserWriter : sirve para devolver la respuesta devuelta al usuario
// http.Request lo que hace es recibir la peticion del usuario.
func ProcesarCrearCuenta(rw http.ResponseWriter, r *http.Request) {
	// Si metodo es diferente a metodo post envie un error
	if r.Method != http.MethodPost{
		// La funcion error recibe un rw(respuesta al usuario), error en string, tipo del error 
		http.Error(rw, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Parcea los datos que vienen en la peticion del usuario
	if err := r.ParseForm(); err != nil {
		http.Error(rw, "Error al leer el formulario", http.StatusBadRequest)
		return	
	}

	// nombre__usuario := r.FormValue("nombre")
	// apellido__usaurio := r.FormValue("apellido")
}