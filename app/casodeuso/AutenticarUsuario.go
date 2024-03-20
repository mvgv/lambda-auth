package casodeuso

import (
	"github.com/mvgv/lambda-auth/app/dominio"
)

type AutenticarUsuario interface {
	AutenticarClienteAnonimo() (string, error)
	AutenticarCliente(cliente *dominio.Cliente) (string, error)
}
