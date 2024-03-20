package casodeuso_test

import (
	"testing"

	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/casodeuso"
	"github.com/mvgv/lambda-auth/app/dominio"
	"github.com/mvgv/lambda-auth/app/repositorio"
	"github.com/stretchr/testify/assert"
)

// Mock do repositório do cliente para os testes

func TestAtualizarCliente_AtualizarCliente(t *testing.T) {
	repositorioMock := repositorio.NewRepositorioClienteMock()
	uc := casodeuso.NewAtualizarClienteImpl(repositorioMock)

	t.Run("AtualizarCliente_Ativo", func(t *testing.T) {
		inputCliente := &dominio.Cliente{
			CPF:    "12345678900",
			ID:     "123",
			Nome:   "John Doe",
			Email:  "john.doe@example.com",
			Status: "ATIVO",
		}

		novosDadosCliente := apresentacao.NewClienteDTO("12345678900", "123", "John Doe", "john.doe@example.com", "ATIVO")

		expectedID := "39b5177e82858ecc5661a2077b58edc3"
		expectedStatus := "ATIVO"

		clienteAtualizado, err := uc.AtualizarCliente(inputCliente, novosDadosCliente)
		assert.NoError(t, err)
		assert.NotNil(t, clienteAtualizado)
		assert.Equal(t, expectedID, clienteAtualizado.ID)
		assert.Equal(t, expectedStatus, clienteAtualizado.Status)
	})

	t.Run("AtualizarCliente_Inativo", func(t *testing.T) {
		inputCliente := &dominio.Cliente{
			CPF:    "12345678900",
			ID:     "123",
			Nome:   "John Doe",
			Email:  "john.doe@example.com",
			Status: "INATIVO",
		}
		novosDadosCliente := apresentacao.NewClienteDTO("12345678900", "123", "John Doe", "john.doe@example.com", "ATIVO")
		expectedID := "123" // MD5 hash of "12345678900"
		expectedStatus := "INATIVO"

		clienteAtualizado, err := uc.AtualizarCliente(inputCliente, novosDadosCliente)
		assert.NoError(t, err)
		assert.NotNil(t, clienteAtualizado)
		assert.Equal(t, expectedID, clienteAtualizado.ID)
		assert.Equal(t, expectedStatus, clienteAtualizado.Status)
	})

}
