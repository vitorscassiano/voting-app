package wires

type PostUserIn struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	CPF      string `json:"cpf" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type PostUserOut struct{}

type GetUserIn struct{}
type GetUserOut struct{}
