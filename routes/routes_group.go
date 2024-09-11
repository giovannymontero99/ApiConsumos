package routes

import (
	"consumosapi/handlers"

	"github.com/gin-gonic/gin"
)

// Manejador de rutas
func RoutesGroup(routes *gin.Engine) {

	// para versiones 1, realizamos el grupo correspondiente.
	v1 := routes.Group("/v1")
	{
		// GET - Pruebas
		v1.GET("/test", handlers.LoginEndPointHandler)
		// POST - login de usuarios
		v1.POST("/login", handlers.LoginEndPointHandler)
	}
}
