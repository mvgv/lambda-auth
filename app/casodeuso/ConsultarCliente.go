package casodeuso

import (
	"github.com/mvgv/lambda-auth/app/dominio"
)

// ConsultarCliente é a interface que define o caso de uso de consulta de cliente
type ConsultarCliente interface {
	ConsultarCliente(idCliente string) (*dominio.Cliente, error)
}
