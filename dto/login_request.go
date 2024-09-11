package dto

// LoginRequest estructura utilizada para recibir los datos de login desde el cliente.
type LoginRequest struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
