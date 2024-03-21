// usecases/autenticacao_cliente_uc.go
package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/mvgv/lambda-auth/app/dominio"
	"golang.org/x/crypto/bcrypt"
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

func (uc *AutenticarUsuarioImpl) ValidarSenha(senhaEntrada string, senhaArmazenada string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(senhaArmazenada), []byte(senhaEntrada))
	if err != nil {
		return false, nil
	}
	return true, nil
}
