package controladores

import (
	"fmt"

	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/casodeuso"
)

type AutenticacaoController struct {
	consultarClienteUC    casodeuso.ConsultarCliente
	autenticacaoClienteUC casodeuso.AutenticarUsuario
}

func NewAutenticacaoController(consultarClienteUC casodeuso.ConsultarCliente, autenticacaoClienteUC casodeuso.AutenticarUsuario) *AutenticacaoController {
	return &AutenticacaoController{
		consultarClienteUC:    consultarClienteUC,
		autenticacaoClienteUC: autenticacaoClienteUC,
	}
}

func (c *AutenticacaoController) Handle(clienteEntrada *apresentacao.ClienteDTO) ([]byte, error) {
	var token string

	cliente, err := c.consultarClienteUC.ConsultarCliente(clienteEntrada.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate client: %v", err)
	}

	senhaValidada, err := c.autenticacaoClienteUC.ValidarSenha(clienteEntrada.Senha, cliente.Senha)
	if err != nil || senhaValidada == false {
		return nil, fmt.Errorf("failed to authenticate client: %v", err)
	}
	token, err = c.autenticacaoClienteUC.AutenticarCliente(cliente)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate client: %v", err)
	}

	return []byte(token), nil
}
