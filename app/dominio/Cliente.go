package dominio

import "golang.org/x/crypto/bcrypt"

type Cliente struct {
	Email  string
	Status string
	Senha  string
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func NewCliente(email, status, senha string) (*Cliente, error) {
	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(senha), 4)

	if err != nil {
		return nil, err
	}

	return &Cliente{
		Email:  email,
		Status: status,
		Senha:  string(hashedSenha),
	}, nil
}
