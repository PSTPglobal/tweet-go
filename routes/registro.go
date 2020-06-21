package routes

import (
	"encoding/json"
	"net/http"
	"github.com/PSTPglobal/tweet-go/bd"
	"github.com/PSTPglobal/tweet-go/models"
)

func Registro(writer http.ResponseWriter, response *http.Request)  {
	var usuario models.Usuario
	err := json.NewDecoder(response.Body).Decode(&usuario)
	if err != nil {
		http.Error(writer, "Error en los datos recibidos " + err.Error(), 400)
		return
	}
	if len( usuario.Email ) == 0 {
		http.Error(writer, "El email de usuario es requerido", 400)
		return
	}
	if len( usuario.Password ) < 6 {
		http.Error(writer, "La contraseña debe ser de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(usuario.Email)
	if encontrado == true {
		http.Error(writer, "Ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(usuario)
	if err != nil {
		http.Error(writer, "Ocurrio un error al intertar realizar el registro de usuario" + err.Error(), 400)
		return
	}
	if status == false {
		http.Error(writer, "No se ha logrado realizar el registro de usuario", 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

