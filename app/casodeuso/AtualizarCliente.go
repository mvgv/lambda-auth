package casodeuso

import (
	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/dominio"
)

type AtualizarCliente interface {
	AtualizarCliente(inputCliente *dominio.Cliente, novosDadosCliente *apresentacao.ClienteDTO) (*dominio.Cliente, error)
}
