package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mvgv/lambda-auth/app/apresentacao"
	"github.com/mvgv/lambda-auth/app/casodeuso"
	"github.com/mvgv/lambda-auth/app/controladores"
	"github.com/mvgv/lambda-auth/app/repositorio"
)

type Response struct {
	Message string `json:"message"`
}

func CustomAuthorizerHandler(ctx context.Context, req events.APIGatewayCustomAuthorizerRequest,
	consultarClienteUC casodeuso.ConsultarCliente, autorizarUsuarioUC casodeuso.AutorizarUsuario) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println("req.AuthorizationToken: ", req.AuthorizationToken)
	token := strings.TrimPrefix(req.AuthorizationToken, "Bearer ")
	respAuthn, err := controladores.NewAutorizarcaoController(consultarClienteUC, autorizarUsuarioUC).Handle(token, req.MethodArn)

	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, fmt.Errorf("failed to parse token: %v", err)
	}

	return respAuthn, nil

}

func AutenticacaoClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, autenticacaoClienteUC casodeuso.AutenticarUsuario,
	consultarClienteUC casodeuso.ConsultarCliente) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de autenticação do cliente
	controller := controladores.NewAutenticacaoController(consultarClienteUC, autenticacaoClienteUC)
	var clienteEntrada *apresentacao.ClienteDTO
	err := json.Unmarshal([]byte(req.Body), &clienteEntrada)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}
	respBody, err := controller.Handle(clienteEntrada)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "mensagem: Funcionario não encontrado"}, fmt.Errorf("failed to handle request: %v", err)
	}
	returnJson, _ := json.Marshal(apresentacao.NewAuthDTO(string(respBody)))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnJson),
	}, nil
}

func CadastroClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, cadastrarClienteUC casodeuso.CadastrarCliente) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de criação de cliente
	controller := controladores.NewCadastroClienteController(cadastrarClienteUC)
	log.Printf("req.Body: %s\n", req.Body)
	respBody, err := controller.Handle(req.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func ConsultaClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest,
	consultarClienteUC casodeuso.ConsultarCliente) (events.APIGatewayProxyResponse, error) {
	controller := controladores.NewConsultaClienteController(consultarClienteUC)
	respBody, err := controller.Handle(req.PathParameters["id_funcionario"])
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "mensagem: Funcionario não encontrado"}, fmt.Errorf("failed to handle request: %v", err)
	}

	returnJson, _ := json.Marshal(apresentacao.NewClienteDTO(respBody.Email,
		respBody.Status, ""))

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnJson),
	}, nil
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest,
	autenticacaoClienteUC casodeuso.AutenticarUsuario,
	consultarClienteUC casodeuso.ConsultarCliente,
	cadastrarClienteUC casodeuso.CadastrarCliente) (events.APIGatewayProxyResponse, error) {

	log.Printf("req.Path: %s\n", req.Path)
	switch req.HTTPMethod {
	case "POST":
		if strings.HasSuffix(req.Path, "/auth") {
			return AutenticacaoClienteHandler(ctx, req, autenticacaoClienteUC, consultarClienteUC)
		} else if req.Path == "/funcionarios" {
			return CadastroClienteHandler(ctx, req, cadastrarClienteUC)
		}
	case "GET":
		return ConsultaClienteHandler(ctx, req, consultarClienteUC)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil
}

func main() {
	clienteRepository := repositorio.NewRepositorioClienteImpl()
	autenticacaoClienteUC := casodeuso.NewAutenticarUsuarioImpl()
	consultarClienteUC := casodeuso.NewConsultarClienteImpl(clienteRepository)
	cadastrarClienteUC := casodeuso.NewCadastrarClienteImpl(clienteRepository)
	autorizarUsuarioUC := casodeuso.NewAutorizarUsuarioImpl()

	lambda.Start(func(ctx context.Context, req map[string]interface{}) (interface{}, error) {
		fmt.Printf("req: %v\n", req)

		// Verificar se é um evento de proxy
		if req["requestContext"] != nil {
			// É um evento de proxy
			proxyRequestJSON, err := json.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("event type not supported")
			}
			var proxyRequestObj events.APIGatewayProxyRequest
			if err := json.Unmarshal(proxyRequestJSON, &proxyRequestObj); err != nil {
				return nil, fmt.Errorf("event type not supported")
			}
			fmt.Printf("proxyRequest: %v\n", proxyRequestObj)
			return Handler(ctx, proxyRequestObj, autenticacaoClienteUC, consultarClienteUC, cadastrarClienteUC)
		}

		// Verificar se é um evento de autorização
		if req["authorizationToken"] != nil {
			// É um evento de autorização
			authorizerRequestJSON, err := json.Marshal(req)
			if err != nil {
				return nil, fmt.Errorf("event type not supported")
			}
			var authorizerRequestObj events.APIGatewayCustomAuthorizerRequest

			if err := json.Unmarshal(authorizerRequestJSON, &authorizerRequestObj); err != nil {
				return nil, fmt.Errorf("event type not supported")
			}
			fmt.Printf("authorizerRequest: %v\n", authorizerRequestObj)

			authorizerResponse, err := CustomAuthorizerHandler(ctx, authorizerRequestObj, consultarClienteUC, autorizarUsuarioUC)
			if err != nil {
				return nil, err
			}
			fmt.Printf("authorizerRequest: %v\n", authorizerRequestObj)
			return authorizerResponse, nil
		}

		// Se não for nem um evento de proxy nem de autorização, retorne um erro
		return nil, fmt.Errorf("event type not supported")
	})

}
