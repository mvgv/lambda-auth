// controllers/CadastroClienteController.go
package controladores

import (
	"encoding/json"
	"fmt"

	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/casodeuso"
	"github.com/mvgv/lambda-auth/app/dominio"
)

type CadastroClienteController struct {
	cadastroClienteUC casodeuso.CadastrarCliente
}

func NewCadastroClienteController(cadastroClienteUC casodeuso.CadastrarCliente) *CadastroClienteController {
	return &CadastroClienteController{
		cadastroClienteUC: cadastroClienteUC,
	}
}

func (controller *CadastroClienteController) Handle(requestBody string) ([]byte, error) {
	var clienteDTO apresentacao.ClienteDTO
	var cliente *dominio.Cliente
	fmt.Printf("req.Body: %s\n", requestBody)

	err := json.Unmarshal([]byte(requestBody), &clienteDTO)

	if err != nil {
		return nil, err
	}

	cliente, err = controller.cadastroClienteUC.CadastrarCliente(clienteDTO)
	if err != nil {
		return nil, err
	}

	response := map[string]string{
		"message": "Cliente cadastrado com sucesso",
		"email":   cliente.Email,
	}
	respBody, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
