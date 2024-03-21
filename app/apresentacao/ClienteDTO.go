package apresentacao

type ClienteDTO struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	Senha  string `json:"senha"`
}

func NewClienteDTO(email, senha, status string) *ClienteDTO {
	return &ClienteDTO{
		Email:  email,
		Status: status,
		Senha:  senha,
	}

}
