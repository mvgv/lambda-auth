// repositories/cliente_repository.go
package repositorio

import (
	"github.com/mvgv/lambda-auth/app/dominio"
)

type RepositorioCliente interface {
	SalvarCliente(cliente *dominio.Cliente) error
	BuscarClientePorID(idCliente string) (*dominio.Cliente, error)
	AtualizarCliente(cliente *dominio.Cliente) error
}
