package wires

type PostAuthentication struct {
	Email    string `json:"email"`
	CPF      string `json:"cpf"`
	Password string `json:"password" binding:"required"`
}
