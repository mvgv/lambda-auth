package casodeuso

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/dominio"
	"github.com/mvgv/lambda-auth/app/repositorio"
)

type AtualizarClienteImpl struct {
	clienteRepository repositorio.RepositorioCliente
}

func NewAtualizarClienteImpl(clienteRepository repositorio.RepositorioCliente) *AtualizarClienteImpl {
	return &AtualizarClienteImpl{
		clienteRepository: clienteRepository,
	}
}

func (uc *AtualizarClienteImpl) AtualizarCliente(inputCliente *dominio.Cliente, novosDadosCliente *apresentacao.ClienteDTO) (*dominio.Cliente, error) {
	var cliente *dominio.Cliente
	var err error
	if novosDadosCliente.Status == "INATIVO" {
		cliente, err = dominio.NewCliente(
			"",
			inputCliente.ID,
			"",
			"",
			"INATIVO",
		)
	} else {
		hash := md5.Sum([]byte(inputCliente.CPF))
		cliente, err = dominio.NewCliente(
			inputCliente.CPF,
			hex.EncodeToString(hash[:]),
			inputCliente.Nome,
			inputCliente.Email,
			"ATIVO",
		)
	}

	if err != nil {
		return nil, err
	}

	err = uc.clienteRepository.AtualizarCliente(cliente)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
