package main

import (
	"consumosapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicio del programa
	router := gin.Default()

	// Maneja las rutas
	routes.RoutesGroup(router)

	// Corre el programa
	router.Run()
}
