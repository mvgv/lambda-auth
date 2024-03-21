package controladores

import (
	"github.com/mvgv/lambda-auth/app/casodeuso"
	"github.com/mvgv/lambda-auth/app/dominio"
)

type ConsultaClienteController struct {
	consultarClienteUC casodeuso.ConsultarCliente
}

func NewConsultaClienteController(consultarClienteUC casodeuso.ConsultarCliente) *ConsultaClienteController {
	return &ConsultaClienteController{consultarClienteUC: consultarClienteUC}
}

func (c *ConsultaClienteController) Handle(idCliente string) (*dominio.Cliente, error) {
	cliente, err := c.consultarClienteUC.ConsultarCliente(idCliente)

	if err != nil {
		return nil, err
	}

	return cliente, nil
}
