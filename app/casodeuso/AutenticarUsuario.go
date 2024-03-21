package casodeuso

import (
	"github.com/mvgv/lambda-auth/app/dominio"
)

type AutenticarUsuario interface {
	AutenticarCliente(cliente *dominio.Cliente) (string, error)
}
