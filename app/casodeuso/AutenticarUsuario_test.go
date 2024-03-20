package casodeuso_test

import (
	"testing"

	"github.com/mvgv/lambda-auth/app/casodeuso"
	"github.com/mvgv/lambda-auth/app/dominio"
	"github.com/stretchr/testify/assert"
)

func TestAutenticarUsuario_AutenticarClienteAnonimo(t *testing.T) {
	uc := casodeuso.NewAutenticarUsuarioImpl()

	t.Run("AutenticarClienteAnonimo", func(t *testing.T) {
		tokenString, err := uc.AutenticarClienteAnonimo()
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)
	})
}

func TestAutenticarUsuario_AutenticarCliente(t *testing.T) {
	uc := casodeuso.NewAutenticarUsuarioImpl()

	t.Run("AutenticarCliente", func(t *testing.T) {
		cliente := &dominio.Cliente{
			ID: "123",
		}
		tokenString, err := uc.AutenticarCliente(cliente)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)
	})

}
