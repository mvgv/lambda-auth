// usecases/autenticacao_cliente_uc.go
package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/mvgv/lambda-auth/app/dominio"
)

type AutenticarUsuarioImpl struct{}

func NewAutenticarUsuarioImpl() *AutenticarUsuarioImpl {
	return &AutenticarUsuarioImpl{}
}

func (uc *AutenticarUsuarioImpl) AutenticarCliente(cliente *dominio.Cliente) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("none"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = cliente.Email
	claims["iss"] = "hackathoncompany.com.br"

	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return "", fmt.Errorf("failed to create unsigned token: %v", err)
	}

	return tokenString, nil
}
