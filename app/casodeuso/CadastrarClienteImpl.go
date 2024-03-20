// casodeuso/CadastrarCliente.go
package casodeuso

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/dominio"
	"github.com/mvgv/lambda-auth/app/repositorio"
)

type CadastrarClienteImpl struct {
	clienteRepository repositorio.RepositorioCliente
}

func NewCadastrarClienteImpl(clienteRepository repositorio.RepositorioCliente) *CadastrarClienteImpl {
	return &CadastrarClienteImpl{
		clienteRepository: clienteRepository,
	}
}

func (uc *CadastrarClienteImpl) CadastrarCliente(inputCliente apresentacao.ClienteDTO) (*dominio.Cliente, error) {
	hash := md5.Sum([]byte(inputCliente.CPF))

	cliente, err := dominio.NewCliente(
		inputCliente.CPF,
		hex.EncodeToString(hash[:]),
		inputCliente.Nome,
		inputCliente.Email,
		"ATIVO",
	)

	if err != nil {
		return nil, err
	}

	err = uc.clienteRepository.SalvarCliente(cliente)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
