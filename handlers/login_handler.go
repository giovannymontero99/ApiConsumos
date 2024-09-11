package handlers

import (
	"consumosapi/config"
	"consumosapi/dto"
	"consumosapi/models"
	"consumosapi/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginEndPointHandler maneja la ruta /login.
// Este manejador recibe las credenciales del usuario, las valida contra la base de datos y devuelve una respuesta adecuada.
func LoginEndPointHandler(c *gin.Context) {

	// Estructura utilizada para recibir los datos de login desde el cliente.
	var loginDataRQ dto.LoginRequest

	// Declara una variable para almacenar los datos del usuario obtenidos de la base de datos.
	var user models.OpeTusers

	// Vincula los datos JSON del cuerpo de la solicitud a la estructura Login.
	if err := c.ShouldBindJSON(&loginDataRQ); err != nil {
		// Si hay un error en la vinculación, responde con un estado HTTP 400 (Bad Request).
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	// Configura la conexión a la base de datos.
	db, err := config.DBConfig()
	if err != nil {
		// Si hay un error al conectarse a la base de datos, responde con un estado HTTP 500 (Internal Server Error).
		c.JSONP(http.StatusInternalServerError, gin.H{"err": "Error al conectarse a la base de datos"})
		return
	}

	result := db.Where(&models.OpeTusers{Identificador: loginDataRQ.User}).Find(&user)

	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		c.JSONP(http.StatusUnauthorized, gin.H{"error": "Usuario no existe"})
		return
	}

	// Estructura utilizada para responder.
	loginResponse := dto.LoginResponse{
		User:    user.Identificador,
		Message: "Buena Esa Calamardo",
	}

	token, err := utils.GenerateToken(user.Identificador, "admin")
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("auth_token", token, -1, "/", "", true, true)
	c.JSONP(200, gin.H{"Usuario": loginResponse, "Token": token})
}
