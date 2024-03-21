package casodeuso

import (
	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/dominio"
)

// CadastrarCliente é a interface que define o caso de uso de cadastro de cliente
type CadastrarCliente interface {
	CadastrarCliente(inputCliente apresentacao.ClienteDTO) (*dominio.Cliente, error)
}
