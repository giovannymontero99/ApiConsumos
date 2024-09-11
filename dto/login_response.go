package dto

// LoginResponse estructura utilizada para enviar la respuesta de login al cliente.
type LoginResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
}
