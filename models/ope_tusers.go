package models

import "gorm.io/gorm"

type OpeTusers struct {
	gorm.ModelUser
	IdUsuario     int    `gorm:"column:usu_usuario;primarykey"`
	Identificador string `gorm:"column:usu_identificador" binding:"required"`
	Email         string `gorm:"column:usu_email"`
	Clave         string `gorm:"column:usu_clave" binding:"required"`
}
