package daemon

import (
	"context"
	"net/http"

	connect "connectrpc.com/connect"
	secretv1 "github.com/codetent/crypta/gen/secret/v1"
	"github.com/codetent/crypta/gen/secret/v1/secretv1connect"
)

var secrets map[string]string = map[string]string{}

type secretServiceServer struct {
	secretv1connect.UnimplementedSecretServiceHandler
}

func NewSecretServiceHandler() (string, http.Handler) {
	return secretv1connect.NewSecretServiceHandler(&secretServiceServer{})
}

func (s *secretServiceServer) SetSecret(
	ctx context.Context,
	req *connect.Request[secretv1.SetSecretRequest],
) (*connect.Response[secretv1.SetSecretResponse], error) {
	name := req.Msg.GetName()
	value := req.Msg.GetValue()
	secrets[name] = value
	return connect.NewResponse(&secretv1.SetSecretResponse{}), nil
}

func (s *secretServiceServer) GetSecret(
	ctx context.Context,
	req *connect.Request[secretv1.GetSecretRequest],
) (*connect.Response[secretv1.GetSecretResponse], error) {
	name := req.Msg.GetName()
	value, exists := secrets[name]
	return connect.NewResponse(&secretv1.GetSecretResponse{
		Value:  value,
		Exists: exists,
	}), nil
}
